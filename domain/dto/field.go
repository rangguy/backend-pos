package dto

import (
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type FieldRequest struct {
	Name         string                 `form:"name" validate:"required"`
	Code         string                 `form:"code" validate:"required"`
	PricePerHour int                    `form:"pricePerHour" validate:"required"`
	Images       []multipart.FileHeader `form:"images" validate:"required"`
}

type UpdateFieldRequest struct {
	Name         string                 `form:"name" validate:"required"`
	Code         string                 `form:"code" validate:"required"`
	PricePerHour int                    `form:"pricePerHour" validate:"required"`
	Images       []multipart.FileHeader `form:"images"`
}

type FieldResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	PricePerHour int        `json:"pricePerHour"`
	Images       []string   `json:"images"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type FieldDetailResponse struct {
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	PricePerHour int        `json:"pricePerHour"`
	Images       []string   `json:"images"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type FieldRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}
