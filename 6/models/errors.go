package models

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrUnAuthorize = errors.New("Unauthorize")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist or duplicate")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Bad Request")

	ErrPublicKey = errors.New("invalid Public Key")
)

//authValidation
var (
	ErrUsername = errors.New("Please Check again your Username")
	ErrPassword = errors.New("Please Check again your Password")
)
