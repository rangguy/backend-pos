package controllers

import (
	errWrap "backend/common/error"
	"backend/common/response"
	"backend/domain/dto"
	"backend/services"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	GetUserLogin(*fiber.Ctx) error
	GetUserByUUID(*fiber.Ctx) error
}

func NewUserController(service services.IServiceRegistry) IUserController {
	return &UserController{service: service}
}

func (u *UserController) Login(ctx *fiber.Ctx) error {
	request := &dto.LoginRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		var syntaxError *json.SyntaxError
		statusCode := http.StatusUnprocessableEntity

		if errors.As(err, &syntaxError) {
			statusCode = http.StatusBadRequest
		}

		errMessage := http.StatusText(statusCode)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    statusCode,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	user, err := u.service.GetUser().Login(ctx.Context(), request)
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusUnprocessableEntity,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user.User,
		Token: &user.Token,
		Fiber: ctx,
	})
}

func (u *UserController) Register(ctx *fiber.Ctx) error {
	request := &dto.RegisterRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		var syntaxError *json.SyntaxError
		statusCode := http.StatusUnprocessableEntity

		if errors.As(err, &syntaxError) {
			statusCode = http.StatusBadRequest
		}

		errMessage := http.StatusText(statusCode)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    statusCode,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	user, err := u.service.GetUser().Register(ctx.Context(), request)
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusUnprocessableEntity,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user.User,
		Fiber: ctx,
	})
}

func (u *UserController) Update(ctx *fiber.Ctx) error {
	request := &dto.UpdateRequest{}
	uuid := ctx.Params("uuid")

	err := ctx.BodyParser(request)
	if err != nil {
		var syntaxError *json.SyntaxError
		statusCode := http.StatusUnprocessableEntity

		if errors.As(err, &syntaxError) {
			statusCode = http.StatusBadRequest
		}

		errMessage := http.StatusText(statusCode)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    statusCode,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	user, err := u.service.GetUser().Update(ctx.Context(), request, uuid)
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusUnprocessableEntity,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user,
		Fiber: ctx,
	})
}

func (u *UserController) GetUserLogin(ctx *fiber.Ctx) error {
	user, err := u.service.GetUser().GetUserLogin(ctx.Context())
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user,
		Fiber: ctx,
	})
}

func (u *UserController) GetUserByUUID(ctx *fiber.Ctx) error {
	user, err := u.service.GetUser().GetUserByUUID(ctx.Context(), ctx.Params("uuid"))
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user,
		Fiber: ctx,
	})
}
