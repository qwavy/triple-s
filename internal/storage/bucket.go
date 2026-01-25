package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"triple-s/internal/models"
)

type BucketStorage struct {
	filePath string
}

func NewBucketStorage(path string) *BucketStorage {
	return &BucketStorage{filePath: path}
}

func (s *BucketStorage) Exists(name string) (bool, error) {
	const op = "storage.bucket.Exists"

	file, err := os.Open(s.filePath)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	for _, eachrecord := range records {
		fmt.Println(eachrecord)
		if eachrecord[0] == name {
			return true, nil
		}
	}

	return false, nil
}

func (s *BucketStorage) Save(b models.Bucket) error {
	const op = "storage.bucket.Save"
	file, err := os.OpenFile("data/Buckets.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{b.Name})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
