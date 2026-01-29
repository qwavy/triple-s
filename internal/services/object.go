package services

import (
	"fmt"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/storage"
)

type ObjectService struct {
	objectStorage *storage.ObjectStorage
	bucketStorage *storage.BucketStorage
}

func NewObjectService(objectStorage *storage.ObjectStorage, bucketStorage *storage.BucketStorage) *ObjectService {
	return &ObjectService{objectStorage, bucketStorage}
}

//func (s *ObjectService) CreateNewObject(bucketName string, objectName string) (newObject, error) {
//	const op = "services.object.CreateNewObject"
//
//	bucketExists, err := s.bucketStorage.Exists(bucketName)
//
//	if !bucketExists {
//		return nil, err
//	}
//
//}

func (s *ObjectService) Delete(bucketName, objectKey string) error {
	const op = "services.object.Delete"

	exists, err := s.objectStorage.Exists(bucketName, objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: %w", op, errors2.ErrBucketNotFound)
	}
	s.objectStorage.Delete(bucketName, objectKey)
	return nil
}
