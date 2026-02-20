package services

import (
	"fmt"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
	"triple-s/internal/validator"
)

func (s *Service) CreateNewBucket(name string) (models.Bucket, error) {
	const op = "services.bucket.CreateNewBucket"

	if !validator.NameRegex.MatchString(name) {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, models.ErrBucketInvalidName)
	}

	exists, err := s.store.IsExistsBucket(name)

	if err != nil {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, models.ErrBucketAlreadyExists)
	}

	now := pkg.GetTime()
	newBucket := models.Bucket{Name: name, CreationDate: now}
	if err := s.store.CreateBucket(newBucket); err != nil {
		return models.Bucket{}, fmt.Errorf("%s: %w", op, err)
	}

	return newBucket, nil
}

func (s *Service) ListBucket() (*models.BucketsResult, error) {
	const op = "services.bucket.List"

	buckets, err := s.store.ListBucket()
	if err != nil {
		return &models.BucketsResult{}, fmt.Errorf("%s: %w", op, err)
	}
	response := &models.BucketsResult{"Nursultan", buckets}
	fmt.Println(response)
	return response, nil
}

func (s *Service) DeleteBucket(name string) error {
	const op = "services.bucket.Delete"

	exists, err := s.store.IsExistsBucket(name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: %w", op, models.ErrBucketNotFound)
	}

	isEmpty, err := s.store.IsEmptyBucket(name)

	if !isEmpty {
		return fmt.Errorf("%s: %w", op, models.ErrBucketNotEmpty)
	}
	err = s.store.DeleteBucket(name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
