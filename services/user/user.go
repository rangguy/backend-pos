package services

import (
	"backend/config"
	"backend/constants"
	errUser "backend/constants/error/user"
	"backend/domain/dto"
	"backend/domain/models"
	"backend/repositories"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserService struct {
	repository repositories.IRepositoryRegistry
}

type IUserService interface {
	Login(context.Context, *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(context.Context, *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Update(context.Context, *dto.UpdateRequest, string) (*dto.UserResponse, error)
	GetUserLogin(context.Context) (*dto.UserResponse, error)
	GetUserByUUID(context.Context, string) (*dto.UserResponse, error)
}

type Claims struct {
	User *dto.UserResponse
	jwt.RegisteredClaims
}

func NewUserService(repository repositories.IRepositoryRegistry) IUserService {
	return &UserService{repository: repository}
}

func (u *UserService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.repository.GetUser().FindByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(time.Duration(config.Config.JwtExpirationTime) * time.Minute).Unix()
	data := &dto.UserResponse{
		UUID:        user.UUID,
		Name:        user.Name,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Role:        strings.ToLower(user.Role.Code),
	}

	claims := &Claims{
		User: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecretKey))
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		User:  *data,
		Token: tokenString,
	}

	return response, nil
}

func (u *UserService) isUsernameExist(ctx context.Context, username string) bool {
	user, err := u.repository.GetUser().FindByUsername(ctx, username)
	if err != nil {
		return false
	}

	if user != nil {
		return true
	}

	return false
}

func (u *UserService) isEmailExist(ctx context.Context, email string) bool {
	user, err := u.repository.GetUser().FindByEmail(ctx, email)
	if err != nil {
		return false
	}

	if user != nil {
		return true
	}

	return false
}

func (u *UserService) Register(ctx context.Context, request *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if u.isUsernameExist(ctx, request.Username) {
		return nil, errUser.ErrUsernameExist
	}

	if u.isEmailExist(ctx, request.Email) {
		return nil, errUser.ErrEmailExist
	}

	newUser := &dto.RegisterRequest{
		Name:        request.Username,
		Username:    request.Username,
		Password:    string(hashedPassword),
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		RoleID:      constants.Admin,
	}

	user, err := u.repository.GetUser().Register(ctx, newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		User: dto.UserResponse{
			UUID:        user.UUID,
			Name:        user.Name,
			Username:    user.Username,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
		},
	}

	return response, nil
}

func (u *UserService) Update(ctx context.Context, request *dto.UpdateRequest, uuid string) (*dto.UserResponse, error) {
	var (
		password                  string
		checkUsername, checkEmail *models.User
		hashedPassword            []byte
		user, userResult          *models.User
		err                       error
		data                      dto.UserResponse
	)

	user, err = u.repository.GetUser().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	isUsernameExist := u.isUsernameExist(ctx, request.Username)
	if isUsernameExist && user.Username != request.Username {
		checkUsername, err = u.repository.GetUser().FindByUsername(ctx, request.Username)
		if err != nil {
			return nil, err
		}

		if checkUsername != nil {
			return nil, errUser.ErrUsernameExist
		}
	}

	isEmailExist := u.isEmailExist(ctx, request.Email)
	if isEmailExist && user.Email != request.Email {
		checkEmail, err = u.repository.GetUser().FindByEmail(ctx, request.Email)
		if err != nil {
			return nil, err
		}

		if checkEmail != nil {
			return nil, errUser.ErrEmailExist
		}
	}

	if request.Password != nil {
		if *request.Password != *request.ConfirmPassword {
			return nil, errUser.ErrPasswordDoesNotMatch
		}

		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		password = string(hashedPassword)
	}

	userResult, err = u.repository.GetUser().Update(ctx, &dto.UpdateRequest{
		Name:        request.Name,
		Username:    request.Username,
		Password:    &password,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}, uuid)
	if err != nil {
		return nil, err
	}

	data = dto.UserResponse{
		UUID:        userResult.UUID,
		Name:        userResult.Name,
		Username:    userResult.Username,
		PhoneNumber: userResult.PhoneNumber,
		Email:       userResult.Email,
	}

	return &data, nil
}

func (u *UserService) GetUserLogin(ctx context.Context) (*dto.UserResponse, error) {
	var (
		userLogin = ctx.Value(constants.UserLogin).(*dto.UserResponse)
		data      dto.UserResponse
	)

	data = dto.UserResponse{
		UUID:        userLogin.UUID,
		Name:        userLogin.Name,
		Username:    userLogin.Username,
		PhoneNumber: userLogin.PhoneNumber,
		Email:       userLogin.Email,
		Role:        userLogin.Role,
	}

	return &data, nil
}

func (u *UserService) GetUserByUUID(ctx context.Context, uuid string) (*dto.UserResponse, error) {
	user, err := u.repository.GetUser().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	data := dto.UserResponse{
		UUID:        user.UUID,
		Name:        user.Name,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}

	return &data, nil
}
