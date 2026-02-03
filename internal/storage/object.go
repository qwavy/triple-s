package storage

import (
	"fmt"
	"net/http"
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

func (s *ObjectStorage) Create(bucketName, objectKey string, content []byte) error {
	const op = "storage.storage.Create"
	err := pkg.OverWriteDataToFile(content, s.filePath+"/"+bucketName+"/"+objectKey)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	contentType := http.DetectContentType(content)

	todayDate := pkg.GetTodayDate()
	err = pkg.WriteDataToCsv([]any{objectKey, contentType, len(content), todayDate}, s.filePath+"/"+bucketName+"/"+"objects.csv")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
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

func (s *ObjectStorage) Get(bucketName, objectKey string) ([]byte, error) {
	const op = "storage.object.Get"

	content, err := pkg.ReadFile(s.filePath + "/" + bucketName + "/" + objectKey)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return content, nil
}

func (s *ObjectStorage) GetFilePath() string {
	return s.filePath
}
