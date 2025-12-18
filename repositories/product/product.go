package repositories

import (
	errWrap "backend/common/error"
	errConstant "backend/constants/error"
	errProduct "backend/constants/error/product"
	"backend/domain/dto"
	"backend/domain/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	FindAllWithPagination(context.Context, *dto.ProductRequestParam) ([]models.Product, int64, error)
	FindAllWithOutPagination(context.Context) ([]models.Product, error)
	FindByUUID(context.Context, string) (*models.Product, error)
	FindByCode(context.Context, string) (*models.Product, error)
	Create(context.Context, *models.Product) (*models.Product, error)
	Update(context.Context, string, *models.Product) (*models.Product, error)
	Delete(context.Context, string) error
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FindAllWithPagination(ctx context.Context, param *dto.ProductRequestParam) ([]models.Product, int64, error) {
	var (
		products []models.Product
		sort     string
		total    int64
	)

	if param.SortColumn != nil {
		sort = fmt.Sprintf("%s %s", *param.SortColumn, *param.SortOrder)
	} else {
		sort = "created_at desc"
	}

	limit := param.Limit
	offset := (param.Page - 1) * limit
	err := p.db.
		WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order(sort).
		Find(&products).
		Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	err = p.db.
		WithContext(ctx).
		Model(&products).
		Count(&total).
		Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return products, total, nil
}

func (p *ProductRepository) FindAllWithOutPagination(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := p.db.
		WithContext(ctx).
		Find(&products).
		Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return products, nil
}

func (p *ProductRepository) FindByUUID(ctx context.Context, uuid string) (*models.Product, error) {
	var product models.Product
	err := p.db.
		WithContext(ctx).
		Where("uuid = ?", uuid).
		First(&product).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errProduct.ErrProductNotFound)
		}

		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &product, nil
}

func (p *ProductRepository) FindByCode(ctx context.Context, code string) (*models.Product, error) {
	var product models.Product
	err := p.db.
		WithContext(ctx).
		Where("code = ?", code).
		First(&product).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errProduct.ErrProductNotFound)
		}

		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &product, nil
}

func (p *ProductRepository) Create(ctx context.Context, req *models.Product) (*models.Product, error) {
	product := models.Product{
		UUID:      uuid.New(),
		Name:      req.Name,
		Code:      req.Code,
		PriceBuy:  req.PriceBuy,
		PriceSale: req.PriceSale,
		Stock:     req.Stock,
		Unit:      req.Unit,
	}

	err := p.db.WithContext(ctx).Create(&product).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &product, nil
}

func (p *ProductRepository) Update(ctx context.Context, uuid string, req *models.Product) (*models.Product, error) {
	product := models.Product{
		Name:      req.Name,
		PriceBuy:  req.PriceBuy,
		PriceSale: req.PriceSale,
		Stock:     req.Stock,
		Unit:      req.Unit,
	}

	err := p.db.WithContext(ctx).Where("uuid = ?", uuid).Updates(&product).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &product, nil
}

func (p *ProductRepository) Delete(ctx context.Context, uuid string) error {
	err := p.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&models.Product{}).Error
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}
