package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
)

type BucketStorage struct {
	filePath            string
	bucketsInfoFilePath string
}

func NewBucketStorage(filePath, bucketsInfoFilePath string) *BucketStorage {
	return &BucketStorage{filePath: filePath, bucketsInfoFilePath: bucketsInfoFilePath}
}

func (s *BucketStorage) Exists(name string) (bool, error) {
	const op = "storage.bucket.Exists"

	file, err := os.Open(s.bucketsInfoFilePath)
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
		fmt.Println(eachrecord[0])
		fmt.Println(name)
		if eachrecord[0] == name {
			return true, nil
		}
	}

	return false, nil
}

func (s *BucketStorage) Create(b models.Bucket) error {
	const op = "storage.bucket.Save"

	err := os.Mkdir(s.filePath+b.Name, 0755)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	pkg.WriteDataToCsv([]string{b.Name}, s.bucketsInfoFilePath)

	return nil
}

func (s *BucketStorage) List() ([]models.Bucket, error) {
	const op = "storage.bucket.List"

	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var buckets []models.Bucket

	for _, eachrecord := range records {
		BucketCreationDate, _ := time.Parse(eachrecord[1], eachrecord[1])
		b := models.Bucket{Name: eachrecord[0], CreationDate: BucketCreationDate}

		buckets = append(buckets, b)
	}

	return buckets, nil
}

func (s *BucketStorage) Delete(bucketName string) error {
	const op = "storage.bucket.Delete"

	_, err := os.Stat(s.filePath + bucketName)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s: %w", op, errors2.ErrBucketNotEmpty)
	}

	err = os.Remove(s.filePath + bucketName)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	//file, err := os.Open(s.filePath)
	//defer file.Close()
	//
	//reader := csv.NewReader(file)
	//
	//records, err := reader.ReadAll()
	//if err != nil {
	//	return fmt.Errorf("%s: %w", op, err)
	//}
	//
	//file, err = os.Create(s.filePath)
	//if err != nil {
	//	return fmt.Errorf("%s: %w", op, err)
	//}
	//defer file.Close()
	//
	//writer := csv.NewWriter(file)
	//
	//for _, record := range records {
	//	if record[0] != bucketName {
	//		err := writer.Write(record)
	//		if err != nil {
	//			return fmt.Errorf("%s: %w", op, err)
	//		}
	//	}
	//
	//}
	//writer.Flush()

	return nil
}
