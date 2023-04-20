package domain

import "errors"

var (
	badNumberError = errors.New("bad phone number")
	internal       = errors.New("internal")
)
