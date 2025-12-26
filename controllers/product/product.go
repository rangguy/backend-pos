package controllers

import (
	errValidation "backend/common/error"
	"backend/common/response"
	errProduct "backend/constants/error/product"
	"backend/domain/dto"
	productService "backend/services"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductController struct {
	service productService.IServiceRegistry
}

type IProductController interface {
	GetAllWithPagination(*fiber.Ctx) error
	GetAllWithoutPagination(*fiber.Ctx) error
	GetByUUID(*fiber.Ctx) error
	GetByCode(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

func NewProductController(service productService.IServiceRegistry) IProductController {
	return &ProductController{service: service}
}

func (p *ProductController) GetAllWithPagination(ctx *fiber.Ctx) error {
	var params dto.ProductRequestParam
	if err := ctx.QueryParser(&params); err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)

		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Fiber:   ctx,
		})
	}

	result, err := p.service.GetProduct().GetAllWithPagination(ctx.Context(), &params)
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) GetAllWithoutPagination(ctx *fiber.Ctx) error {
	result, err := p.service.GetProduct().GetAllWithoutPagination(ctx.Context())
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) GetByUUID(ctx *fiber.Ctx) error {
	result, err := p.service.GetProduct().GetByUUID(ctx.Context(), ctx.Params("uuid"))
	if err != nil {
		if errors.Is(err, errProduct.ErrProductNotFound) {
			return response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusNotFound,
				Err:   err,
				Fiber: ctx,
			})
		}

		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) GetByCode(ctx *fiber.Ctx) error {
	result, err := p.service.GetProduct().GetByCode(ctx.Context(), ctx.Params("code"))
	if err != nil {
		if errors.Is(err, errProduct.ErrProductNotFound) {
			return response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusNotFound,
				Err:   err,
				Fiber: ctx,
			})
		}

		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) Create(ctx *fiber.Ctx) error {
	request := &dto.ProductRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		var syntaxError *json.SyntaxError
		statusCode := http.StatusUnprocessableEntity

		if errors.As(err, &syntaxError) {
			statusCode = http.StatusBadRequest
		}

		errMessage := http.StatusText(statusCode)
		errResponse := errValidation.ErrValidationResponse(err)

		return response.HttpResponse(response.ParamHTTPResp{
			Code:    statusCode,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)

		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Fiber:   ctx,
		})
	}

	result, err := p.service.GetProduct().Create(ctx.Context(), request)
	if err != nil {
		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) Update(ctx *fiber.Ctx) error {
	request := &dto.UpdateProductRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		var syntaxError *json.SyntaxError
		statusCode := http.StatusUnprocessableEntity

		if errors.As(err, &syntaxError) {
			statusCode = http.StatusBadRequest
		}

		errMessage := http.StatusText(statusCode)
		errResponse := errValidation.ErrValidationResponse(err)

		return response.HttpResponse(response.ParamHTTPResp{
			Code:    statusCode,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Fiber:   ctx,
		})
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)

		return response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     err,
			Message: &errMessage,
			Data:    errorResponse,
			Fiber:   ctx,
		})
	}

	result, err := p.service.GetProduct().Update(
		ctx.Context(),
		ctx.Params("uuid"),
		request,
	)
	if err != nil {
		if errors.Is(err, errProduct.ErrProductNotFound) {
			return response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusNotFound,
				Err:   err,
				Fiber: ctx,
			})
		}

		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  result,
		Fiber: ctx,
	})
}

func (p *ProductController) Delete(ctx *fiber.Ctx) error {
	err := p.service.GetProduct().Delete(ctx.Context(), ctx.Params("uuid"))
	if err != nil {
		if errors.Is(err, errProduct.ErrProductNotFound) {
			return response.HttpResponse(response.ParamHTTPResp{
				Code:  http.StatusNotFound,
				Err:   err,
				Fiber: ctx,
			})
		}

		return response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Err:   err,
			Fiber: ctx,
		})
	}

	return response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Fiber: ctx,
	})
}
