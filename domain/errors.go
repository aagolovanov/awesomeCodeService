package domain

import "errors"

var (
	badNumberError        = errors.New("bad phone number")
	internal              = errors.New("internal") // specific
	badUUIDError          = errors.New("bad requestId sent")
	requestNotExistError  = errors.New("request does not exist")
	attemptsExceededError = errors.New("limit") // specific
	invalidCodeError      = errors.New("invalid code")
)
