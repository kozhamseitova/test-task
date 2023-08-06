package utils

import "errors"

var (
	ErrInternalError = errors.New("Internal Server Error")
	ErrNotFound = errors.New("User Not Found Error")
	ErrUserAlreadyExists = errors.New("User Is Already Exists")
	ErrHeaderIsNotSet = errors.New("Authorization Header Is Not Set")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
) 
