package services

import (
	"fmt"
	"regexp"
	"time"
	"triple-s/internal/errors"
	"triple-s/internal/models"
	"triple-s/internal/storage"
)

type ListAllMyBucketsResult struct {
	Owner   string
	Buckets []models.Bucket `xml:"Buckets>Bucket"`
}

type BucketService struct {
	storage *storage.BucketStorage
}

var nameRegex = regexp.MustCompile(`^[a-z0-9.-]{3,63}$`)

func NewBucketService(bucketStorage *storage.BucketStorage) *BucketService {
	return &BucketService{storage: bucketStorage}
}

func (s *BucketService) CreateNewBucket(name string) (models.Bucket, error) {
	const op = "services.bucket.CreateNewBucket"

	if !nameRegex.MatchString(name) {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, errors.ErrBucketInvalidName)
	}

	exists, err := s.storage.Exists(name)

	if err != nil {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, errors.ErrBucketAlreadyExists)
	}

	now := time.Now()
	newBucket := models.Bucket{Name: name, CreationDate: now}

	if err := s.storage.Save(newBucket); err != nil {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, err)
	}

	return newBucket, nil
}

func (s *BucketService) List() (ListAllMyBucketsResult, error) {
	const op = "services.bucket.List"

	buckets, err := s.storage.List()
	if err != nil {
		return ListAllMyBucketsResult{}, fmt.Errorf("%s: %w", op, err)
	}
	response := ListAllMyBucketsResult{"Nursultan", buckets}

	return response, nil
}
