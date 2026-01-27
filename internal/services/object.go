package services

import "triple-s/internal/storage"

type ObjectService struct {
	objectStorage *storage.ObjectStorage
	bucketStorage *storage.BucketStorage
}

func NewObjectService(objectStorage *storage.ObjectStorage, bucketStorage *storage.BucketStorage) *ObjectService {
	return &ObjectService{objectStorage, bucketStorage}
}

func (s *ObjectService) CreateNewObject(bucketName string, objectName string) (newObject, error) {
	const op = "services.object.CreateNewObject"

	bucketExists, err := s.bucketStorage.Exists(bucketName)

	if !bucketExists {
		return nil, err
	}

}
