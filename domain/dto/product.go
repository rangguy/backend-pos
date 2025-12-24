package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProductRequest struct {
	Code      *string `json:"code"`
	Name      string  `json:"name"  validate:"required"`
	PriceBuy  uint    `json:"price_buy" validate:"required"`
	PriceSale uint    `json:"price_sale" validate:"required"`
	Stock     uint    `json:"stock" validate:"required"`
	Unit      string  `json:"unit" validate:"required"`
}

type UpdateProductRequest struct {
	Code      *string `json:"code"`
	Name      string  `json:"name"`
	PriceBuy  uint    `json:"price_buy"`
	PriceSale uint    `json:"price_sale"`
	Stock     uint    `json:"stock"`
	Unit      string  `json:"unit"`
}

type ProductResponse struct {
	UUID      uuid.UUID  `json:"uuid"`
	Code      *string    `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"price_buy"`
	PriceSale uint       `json:"price_sale"`
	Stock     uint       `json:"stock"`
	Unit      string     `json:"unit"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ProductDetailResponse struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"price_buy"`
	PriceSale uint       `json:"price_sale"`
	Stock     uint       `json:"stock"`
	Unit      string     `json:"unit"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ProductRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}
