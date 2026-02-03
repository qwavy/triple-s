package errors

import "errors"

var (
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrBucketInvalidName   = errors.New("invalid bucket name")

	ErrBucketNotFound = errors.New("bucket not found")
	ErrBucketNotEmpty = errors.New("bucket not empty")

	ErrObjNotFound   = errors.New("object not found")
	ErrObjInvalidKey = errors.New("invalid object key")
)
