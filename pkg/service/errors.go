package service

import "errors"

var (
	ErrorEmptyCountry = errors.New("Empty country response")
	ErrorEmptyName    = errors.New("Empty name input")
)
