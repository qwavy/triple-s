package repository

import (
	"fmt"
	"net/http"
	"os"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
)

func (s *Store) CreateObject(bucketName, objectKey string, content []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.repository.Create"
	err := pkg.OverWriteDataToFile(content, s.filePath+"/"+bucketName+"/"+objectKey)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	contentType := http.DetectContentType(content)

	todayDate := pkg.GetTime()
	err = pkg.WriteDataToCsv([]any{objectKey, contentType, len(content), todayDate}, s.filePath+"/"+bucketName+"/"+"objects.csv")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if err := s.LastModificationTime(bucketName); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Store) IsExistsObject(bucketName, objectKey string) (bool, error) {
	const op = "repository.object.Exists"

	bucketExists := pkg.FolderExists(s.filePath + "/" + bucketName)
	if !bucketExists {
		return false, models.ErrBucketNotFound
	}

	objectExists := pkg.FileExists(s.filePath + "/" + bucketName + "/" + objectKey)
	if !objectExists {
		return false, models.ErrObjNotFound
	}

	return true, nil
}

func (s *Store) DeleteObject(bucketName, objectKey string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.object.Delete"

	err := os.Remove(s.filePath + "/" + bucketName + "/" + objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if err := s.LastModificationTime(bucketName); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Store) GetObject(bucketName, objectKey string) ([]byte, error) {

	const op = "repository.object.Get"

	content, err := pkg.ReadFile(s.filePath + "/" + bucketName + "/" + objectKey)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return content, nil
}
