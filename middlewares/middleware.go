package middlewares

import (
	"backend/common/response"
	"backend/config"
	"backend/constants"
	errConstant "backend/constants/error"
	services "backend/services/user"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func HandlePanic() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := errConstant.ErrInternalServerError.Error()

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		logrus.Errorf("Error occurred: %v\nStack trace: %s", err, debug.Stack())

		return c.Status(code).JSON(response.Response{
			Status:  constants.Error,
			Message: message,
		})
	}
}

func RateLimiter(max int, duration time.Duration) fiber.Handler {
	configLimiter := limiter.Config{
		Max:        max,
		Expiration: duration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(response.Response{
				Status:  constants.Error,
				Message: errConstant.ErrTooManyRequests.Error(),
			})
		},
	}
	return limiter.New(configLimiter)
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func responseUnauthorized(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(response.Response{
		Status:  constants.Error,
		Message: message,
	})
}

func validateAPIKey(c *fiber.Ctx) error {
	apiKey := c.Get(constants.XApiKey)
	requestAt := c.Get(constants.XRequestAt)
	serviceName := c.Get(constants.XServiceName)
	signatureKey := config.Config.SignatureKey

	validateKey := fmt.Sprintf("%s:%s:%s", serviceName, signatureKey, requestAt)
	hash := sha256.New()
	hash.Write([]byte(validateKey))
	resultHash := hex.EncodeToString(hash.Sum(nil))

	if apiKey != resultHash {
		return errConstant.ErrUnauthorized
	}
	return nil
}

func validateBearerToken(c *fiber.Ctx, token string) error {
	if !strings.Contains(token, "Bearer") {
		return errConstant.ErrUnauthorized
	}

	tokenString := extractBearerToken(token)
	if tokenString == "" {
		return errConstant.ErrUnauthorized
	}

	claims := &services.Claims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errConstant.ErrInvalidToken
		}

		jwtSecret := []byte(config.Config.JwtSecretKey)
		return jwtSecret, nil
	})

	if err != nil || !tokenJwt.Valid {
		return errConstant.ErrUnauthorized
	}

	c.Locals(constants.UserLogin, claims.User)
	c.Set(constants.Token, token)
	return nil
}

func CheckRole(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		// Bersihkan prefix "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format. Use: Bearer <token>",
			})
		}

		// Parse dan validasi JWT token
		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// CRITICAL: Validasi algoritma signing untuk mencegah algorithm confusion attack
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Ambil JWT secret dari config
			return []byte(config.Config.JwtSecretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		if !parsedToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is not valid",
			})
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Token has expired",
				})
			}
		}

		userRole, ok := claims["Role"].(string)
		if !ok || userRole == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Role not found in token",
			})
		}

		userRole = strings.ToLower(strings.TrimSpace(userRole))

		hasPermission := false
		for _, allowedRole := range allowedRoles {
			if strings.ToLower(allowedRole) == userRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden: insufficient permissions",
				"message": fmt.Sprintf("Required role: %v, your role: %s", allowedRoles, userRole),
			})
		}

		c.Locals("uuid", claims["UUID"].(string))
		c.Locals("username", claims["Username"].(string))
		c.Locals("role", userRole)

		return c.Next()
	}
}

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		token := c.Get(constants.Authorization)
		if token == "" {
			return responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
		}

		err = validateBearerToken(c, token)
		if err != nil {
			return responseUnauthorized(c, err.Error())
		}

		err = validateAPIKey(c)
		if err != nil {
			return responseUnauthorized(c, err.Error())
		}

		return c.Next()
	}
}
