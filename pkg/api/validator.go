package api

import "gopkg.in/go-playground/validator.v9"

type RequestValidator struct {
	Validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
