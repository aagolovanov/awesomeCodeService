package domain

import "errors"

var (
	BadNumberError        = errors.New("bad phone number")
	Internal              = errors.New("internal") // specific
	BadUUIDError          = errors.New("bad requestId sent")
	RequestNotExistError  = errors.New("request does not exist")
	AttemptsExceededError = errors.New("limit") // specific
	InvalidCodeError      = errors.New("invalid code")
)
