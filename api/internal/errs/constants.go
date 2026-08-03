package errs

import "errors"

var (
	ErrAuthorAlreadyExists   = errors.New("author with that name already exists")
	ErrAccountAlreadyExists  = errors.New("account already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrInvalidCredentials    = errors.New("email or password for the user is incorrect")
	ErrSessionAlreadyExists  = errors.New("session already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrUserNotFound          = errors.New("user with that email or password doesn't exists")
	ErrUnauthorized          = errors.New("not authorized to access this content")
)
