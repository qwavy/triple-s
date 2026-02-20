package models

import (
	"errors"
	"time"
)

var (
	ErrObjNotFound   = errors.New("object not found")
	ErrObjInvalidKey = errors.New("invalid object key")
)

type Object struct {
	ObjectKey    string    `xml:"ObjectKey"`
	Size         int       `xml:"Size"`
	ContentType  string    `xml:"ContentType"`
	LastModified time.Time `xml:"CreationDate"`
}
