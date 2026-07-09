package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amount cloud not be zero")
	ErrRepository = errors.New("repository error")
)
