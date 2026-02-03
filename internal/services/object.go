package services

import (
	"fmt"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/pkg"
	"triple-s/internal/storage"
	"triple-s/internal/validator"
)

type ObjectService struct {
	objectStorage *storage.ObjectStorage
	bucketStorage *storage.BucketStorage
}

func NewObjectService(objectStorage *storage.ObjectStorage, bucketStorage *storage.BucketStorage) *ObjectService {
	return &ObjectService{objectStorage, bucketStorage}
}

func (s *ObjectService) Create(bucketName, objectKey string, content []byte) error {
	const op = "services.object.CreateNewObject"

	bucketExists, err := s.bucketStorage.Exists(bucketName)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !bucketExists {
		return fmt.Errorf("%s: %w", op, errors2.ErrBucketNotFound)
	}

	if !validator.NameRegex.MatchString(objectKey) {
		return fmt.Errorf("%s: %w", op, errors2.ErrBucketInvalidName)
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.objectStorage.Create(bucketName, objectKey, content)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *ObjectService) Delete(bucketName, objectKey string) error {
	const op = "services.object.Delete"

	exists, err := s.objectStorage.Exists(bucketName, objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: %w", op, errors2.ErrBucketNotFound)
	}

	err = s.objectStorage.Delete(bucketName, objectKey)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *ObjectService) Get(bucketName, objectKey string) ([]byte, string, error) {
	const op = "services.object.Get"
	bucketExists, err := s.bucketStorage.Exists(bucketName)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	if !bucketExists {
		return nil, "", fmt.Errorf("%s: %w", op, errors2.ErrBucketNotFound)
	}

	objectExists, err := s.objectStorage.Exists(bucketName, objectKey)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	if !objectExists {
		return nil, "", fmt.Errorf("%s: %w", op, errors2.ErrObjNotFound)
	}

	content, err := s.objectStorage.Get(bucketName, objectKey)

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	contentType, err := pkg.FindRowCsvByName(objectKey, 0, s.objectStorage.GetFilePath()+"/"+bucketName+"/"+"objects.csv")

	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}
	return content, contentType[1], nil
}
