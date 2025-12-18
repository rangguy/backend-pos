package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type ProductRequest struct {
	Name         string                 `form:"name" validate:"required"`
	PricePerHour int                    `form:"pricePerHour" validate:"required"`
	Images       []multipart.FileHeader `form:"images" validate:"required"`
}

type UpdateProductRequest struct {
	Name         string                 `form:"name" validate:"required"`
	PricePerHour int                    `form:"pricePerHour" validate:"required"`
	Images       []multipart.FileHeader `form:"images"`
}

type ProductResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	Name         string     `json:"name"`
	PricePerHour int        `json:"pricePerHour"`
	Images       []string   `json:"images"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type ProductDetailResponse struct {
	Name         string     `json:"name"`
	PricePerHour int        `json:"pricePerHour"`
	Images       []string   `json:"images"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type ProductRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}
