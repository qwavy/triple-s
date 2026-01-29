package storage

import (
	"fmt"
	"os"
	"triple-s/internal/errors"
	"triple-s/internal/pkg"
)

type ObjectStorage struct {
	filePath string
}

func NewObjectStorage(path string) *ObjectStorage {
	return &ObjectStorage{filePath: path}
}

func (s *ObjectStorage) Exists(bucketName, objectKey string) (bool, error) {
	const op = "storage.object.Exists"

	bucketExists := pkg.FolderExists(s.filePath + "/" + bucketName)
	if !bucketExists {
		return false, errors.ErrBucketNotFound
	}

	objectExists := pkg.FileExists(s.filePath + "/" + bucketName + "/" + objectKey)
	if !objectExists {
		return false, errors.ErrObjNotFound
	}

	return true, nil
}

func (s *ObjectStorage) Delete(bucketName, objectKey string) error {
	const op = "storage.object.Delete"

	err := os.Remove(s.filePath + "/" + bucketName + "/" + objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
