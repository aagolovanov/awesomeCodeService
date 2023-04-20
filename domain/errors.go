package domain

import "errors"

var (
	badNumberError       = errors.New("bad phone number")
	internal             = errors.New("internal")
	badUUIDError         = errors.New("bad requestId sent")
	requestNotExistError = errors.New("request does not exist")
)
