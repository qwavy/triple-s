package models

import "errors"

var (
	ErrObjNotFound   = errors.New("object not found")
	ErrObjInvalidKey = errors.New("invalid object key")
)

type Object struct {
}
