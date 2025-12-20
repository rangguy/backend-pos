package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProductRequest struct {
	Code      string `form:"code" validate:"required"`
	Name      string `form:"name" validate:"required"`
	PriceBuy  uint   `form:"price_buy" validate:"required"`
	PriceSale uint   `form:"price_sale" validate:"required"`
	Stock     uint   `form:"stock" validate:"required"`
	Unit      string `form:"unit" validate:"required"`
}

type UpdateProductRequest struct {
	Code      string `form:"code"`
	Name      string `form:"name"`
	PriceBuy  uint   `form:"price_buy"`
	PriceSale uint   `form:"price_sale"`
	Stock     uint   `form:"stock"`
	Unit      string `form:"unit"`
}

type ProductResponse struct {
	UUID      uuid.UUID  `json:"uuid"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"price_buy"`
	PriceSale uint       `json:"price_sale"`
	Stock     uint       `json:"stock"`
	Unit      string     `json:"unit"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ProductDetailResponse struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"price_buy"`
	PriceSale uint       `json:"price_sale"`
	Stock     uint       `json:"stock"`
	Unit      string     `json:"unit"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ProductRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}
