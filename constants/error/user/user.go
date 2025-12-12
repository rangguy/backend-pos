package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrUsernameExist        = errors.New("username exist")
	ErrEmailExist           = errors.New("email exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrEmailExist,
	ErrPasswordDoesNotMatch,
}
