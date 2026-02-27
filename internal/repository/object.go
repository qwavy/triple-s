package repository

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"triple-s/internal/pkg"
)

func (s *Store) CreateObject(bucketName, objectKey string, content []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.object.Create"

	err := pkg.OverWriteDataToFile(content, s.filePath+"/"+bucketName+"/"+objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	contentType := http.DetectContentType(content)
	todayDate := pkg.GetTime()
	objectsCsvPath := s.filePath + "/" + bucketName + "/objects.csv"

	var records [][]string
	if file, err := os.Open(objectsCsvPath); err == nil {
		reader := csv.NewReader(file)
		records, _ = reader.ReadAll()
		file.Close()
	}

	tmpFile, err := os.Create(objectsCsvPath + ".tmp")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	writer := csv.NewWriter(tmpFile)

	found := false
	newRecord := []string{objectKey, contentType, fmt.Sprintf("%d", len(content)), todayDate}

	for _, record := range records {
		if record[0] == objectKey {
			writer.Write(newRecord)
			found = true
		} else {
			writer.Write(record)
		}
	}

	if !found {
		writer.Write(newRecord)
	}

	writer.Flush()
	tmpFile.Close()
	os.Rename(objectsCsvPath+".tmp", objectsCsvPath)

	if err := s.LastModificationTime(bucketName); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Store) DeleteObject(bucketName, objectKey string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.object.Delete"

	err := os.Remove(s.filePath + "/" + bucketName + "/" + objectKey)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	objectsCsvPath := s.filePath + "/" + bucketName + "/objects.csv"
	if file, err := os.Open(objectsCsvPath); err == nil {
		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()
		file.Close()

		tmpFile, _ := os.Create(objectsCsvPath + ".tmp")
		writer := csv.NewWriter(tmpFile)
		for _, record := range records {
			if record[0] != objectKey {
				writer.Write(record)
			}
		}
		writer.Flush()
		tmpFile.Close()
		os.Rename(objectsCsvPath+".tmp", objectsCsvPath)
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

func (s *Store) IsExistsObject(bucketName, objectKey string) (bool, error) {

	bucketExists := pkg.FolderExists(s.filePath + "/" + bucketName)
	if !bucketExists {
		return false, nil
	}

	objectExists := pkg.FileExists(s.filePath + "/" + bucketName + "/" + objectKey)
	if !objectExists {
		return false, nil
	}

	return true, nil
}
