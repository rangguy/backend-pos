package response

import (
	"backend/constants"
	errConstant "backend/constants/error"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Fiber   *fiber.Ctx
	Data    interface{}
	Token   *string
}

func HttpResponse(param ParamHTTPResp) error {
	if param.Err == nil {
		return param.Fiber.Status(param.Code).JSON(Response{
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   param.Token,
		})
	}

	message := errConstant.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		if errConstant.ErrMapping(param.Err) {
			message = param.Err.Error()
		}
	}

	return param.Fiber.Status(param.Code).JSON(Response{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})
}
