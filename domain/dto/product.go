package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProductRequest struct {
	Code      string `form:"code"`
	Name      string `form:"name" validate:"required"`
	PriceBuy  uint   `form:"priceBuy" validate:"required"`
	PriceSale uint   `form:"priceSale" validate:"required"`
	Stock     uint   `form:"stock" validate:"required"`
	Unit      string `form:"unit" validate:"required"`
}

type UpdateProductRequest struct {
	Code      string `form:"code"`
	Name      string `form:"name"`
	PriceBuy  uint   `form:"priceBuy"`
	PriceSale uint   `form:"priceSale"`
	Stock     uint   `form:"stock"`
	Unit      string `form:"unit"`
}

type ProductResponse struct {
	UUID      uuid.UUID  `json:"uuid"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"priceBuy"`
	PriceSale uint       `json:"priceSale"`
	Stock     uint       `json:"stock"`
	Unit      string     `json:"unit"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ProductDetailResponse struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	PriceBuy  uint       `json:"priceBuy"`
	PriceSale uint       `json:"priceSale"`
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
