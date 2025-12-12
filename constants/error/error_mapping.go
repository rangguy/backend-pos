package error

import (
	errProduct "backend/constants/error/product"
	errUser "backend/constants/error/user"
)

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(append(GeneralErrors[:], errUser.UserErrors[:]...), errProduct.ProductErrors[:]...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
