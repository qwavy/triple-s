package models

import (
	"errors"
)

var (
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrBucketInvalidName   = errors.New("invalid bucket name")
	ErrBucketNotFound      = errors.New("bucket not found")
	ErrBucketNotEmpty      = errors.New("bucket not empty")
)

type Bucket struct {
	Name         string `xml:"Name"`
	CreationDate string `xml:"CreationDate"`
}

type ListAllMyBucketsResult struct {
	Owner   string
	Buckets []Bucket `xml:"Buckets>Bucket"`
}
