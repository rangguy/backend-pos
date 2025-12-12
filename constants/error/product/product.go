package error

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductIsExist  = errors.New("product already exist")
)

var ProductErrors = []error{
	ErrProductNotFound,
	ErrProductIsExist,
}
