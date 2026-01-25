package errors

import "errors"

var (
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrBucketInvalidName   = errors.New("bucket have invalid name")
)
