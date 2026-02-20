package services

import (
	"fmt"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
	"triple-s/internal/validator"
)

func (s *Service) CreateObject(bucketName, objectKey string, content []byte) error {
	const op = "services.object.CreateNewObject"

	bucketExists, err := s.store.IsExistsBucket(bucketName)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !bucketExists {
		return fmt.Errorf("%s: %w", op, models.ErrBucketNotFound)
	}

	err = validator.ValidateName(objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, models.ErrBucketInvalidName)
	}

	err = s.store.CreateObject(bucketName, objectKey, content)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) DeleteObject(bucketName, objectKey string) error {
	const op = "services.object.Delete"

	exists, err := s.store.IsExistsObject(bucketName, objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: %w", op, models.ErrBucketNotFound)
	}

	err = s.store.DeleteObject(bucketName, objectKey)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetObject(bucketName, objectKey string) ([]byte, string, error) {
	const op = "services.object.Get"
	bucketExists, err := s.store.IsExistsBucket(bucketName)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	if !bucketExists {
		return nil, "", fmt.Errorf("%s: %w", op, models.ErrBucketNotFound)
	}

	objectExists, err := s.store.IsExistsObject(bucketName, objectKey)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	if !objectExists {
		return nil, "", fmt.Errorf("%s: %w", op, models.ErrObjNotFound)
	}

	content, err := s.store.GetObject(bucketName, objectKey)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	contentType, err := pkg.FindRowCsvByName(objectKey, 0, s.store.GetFilePath()+"/"+bucketName+"/"+"objects.csv")

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}
	return content, contentType[1], nil
}
