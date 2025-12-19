package services

import (
	"backend/common/util"
	"backend/domain/dto"
	"backend/repositories"
	"context"
	uuid2 "github.com/google/uuid"
)

type ProductService struct {
	repository repositories.IRepositoryRegistry
}

type IProductService interface {
	GetAllWithPagination(context.Context, *dto.ProductRequestParam) (*util.PaginationResult, error)
	GetAllWithoutPagination(context.Context) ([]dto.ProductResponse, error)
	GetByUUID(context.Context, string) (*dto.ProductResponse, error)
	GetByCode(context.Context, string) (*dto.ProductResponse, error)
	Create(context.Context, *dto.ProductRequest) (*dto.ProductResponse, error)
	Update(context.Context, string, *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	Delete(context.Context, string) error
}

func NewProductService(repository repositories.IRepositoryRegistry) IProductService {
	return &ProductService{repository: repository}
}

func (p *ProductService) GetAllWithPagination(ctx context.Context, param *dto.ProductRequestParam) (*util.PaginationResult, error) {
	products, total, err := p.repository.GetProduct().FindAllWithPagination(ctx, param)

	if err != nil {
		return nil, err
	}

	productResult := make([]*dto.ProductResponse, 0, len(products))
	for _, product := range products {
		productResult = append(productResult, &dto.ProductResponse{
			UUID:      product.UUID,
			Code:      product.Code,
			Name:      product.Name,
			PriceBuy:  product.PriceBuy,
			PriceSale: product.PriceSale,
			Stock:     product.Stock,
			Unit:      product.Unit,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	pagination := &util.PaginationParam{
		Count: total,
		Page:  param.Page,
		Limit: param.Limit,
		Data:  productResult,
	}

	response := util.GeneratePagination(*pagination)
	return &response, nil
}

func (p *ProductService) GetAllWithoutPagination(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := p.repository.GetProduct().FindAllWithoutPagination(ctx)
	if err != nil {
		return nil, err
	}

	productResult := make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		productResult = append(productResult, dto.ProductResponse{
			UUID:      product.UUID,
			Code:      product.Code,
			Name:      product.Name,
			PriceBuy:  product.PriceBuy,
			PriceSale: product.PriceSale,
			Stock:     product.Stock,
			Unit:      product.Unit,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	return productResult, nil
}

func (p *ProductService) GetByUUID(ctx context.Context, uuid string) (*dto.ProductResponse, error) {
	product, err := p.repository.GetProduct().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	productResult := &dto.ProductResponse{
		UUID:      product.UUID,
		Code:      product.Code,
		Name:      product.Name,
		PriceBuy:  product.PriceBuy,
		PriceSale: product.PriceSale,
		Stock:     product.Stock,
		Unit:      product.Unit,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return productResult, nil
}

func (p *ProductService) GetByCode(ctx context.Context, code string) (*dto.ProductResponse, error) {
	product, err := p.repository.GetProduct().FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	productResult := &dto.ProductResponse{
		UUID:      product.UUID,
		Code:      product.Code,
		Name:      product.Name,
		PriceBuy:  product.PriceBuy,
		PriceSale: product.PriceSale,
		Stock:     product.Stock,
		Unit:      product.Unit,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return productResult, nil
}

func (p *ProductService) Create(ctx context.Context, request *dto.ProductRequest) (*dto.ProductResponse, error) {
	product := &dto.ProductRequest{
		Code:      request.Code,
		Name:      request.Name,
		PriceBuy:  request.PriceBuy,
		PriceSale: request.PriceSale,
		Stock:     request.Stock,
		Unit:      request.Unit,
	}

	newProduct, err := p.repository.GetProduct().Create(ctx, product)
	if err != nil {
		return nil, err
	}

	productResult := &dto.ProductResponse{
		UUID:      newProduct.UUID,
		Code:      newProduct.Code,
		Name:      newProduct.Name,
		PriceBuy:  newProduct.PriceBuy,
		PriceSale: newProduct.PriceSale,
		Stock:     newProduct.Stock,
		Unit:      request.Unit,
		CreatedAt: newProduct.CreatedAt,
		UpdatedAt: newProduct.UpdatedAt,
	}

	return productResult, nil
}

func (p *ProductService) Update(ctx context.Context, uuid string, request *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	updateProduct := &dto.ProductRequest{
		Code:      request.Code,
		Name:      request.Name,
		PriceBuy:  request.PriceBuy,
		PriceSale: request.PriceSale,
		Stock:     request.Stock,
		Unit:      request.Unit,
	}

	newProduct, err := p.repository.GetProduct().Update(ctx, uuid, updateProduct)
	if err != nil {
		return nil, err
	}

	uuidParsed, _ := uuid2.Parse(uuid)

	productResult := &dto.ProductResponse{
		UUID:      uuidParsed,
		Code:      newProduct.Code,
		Name:      newProduct.Name,
		PriceBuy:  newProduct.PriceBuy,
		PriceSale: newProduct.PriceSale,
		Stock:     newProduct.Stock,
		Unit:      newProduct.Unit,
		CreatedAt: newProduct.CreatedAt,
		UpdatedAt: newProduct.UpdatedAt,
	}

	return productResult, nil
}

func (p *ProductService) Delete(ctx context.Context, uuid string) error {
	_, err := p.repository.GetProduct().FindByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	err = p.repository.GetProduct().Delete(ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}
