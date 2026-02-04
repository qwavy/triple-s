package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
)

func (s *Store) IsExistsBucket(name string) (bool, error) {
	const op = "repository.bucket.Exists"

	file, err := os.Open(s.filePath + s.bucketsInfoFilePath)
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
		if eachrecord[0] == name {
			return true, nil
		}
	}

	return false, nil
}

func (s *Store) CreateBucket(b models.Bucket) error {
	const op = "repository.bucket.Save"

	err := os.Mkdir(s.filePath+b.Name, 0755)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	creationTime := time.Now()

	pkg.WriteDataToCsv([]any{b.Name, creationTime}, s.bucketsInfoFilePath)

	return nil
}

func (s *Store) ListBucket() ([]models.Bucket, error) {
	const op = "repository.bucket.List"

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
		b := models.Bucket{Name: eachrecord[0], CreationDate: eachrecord[1]}

		buckets = append(buckets, b)
	}

	return buckets, nil
}

func (s *Store) DeleteBucket(bucketName string) error {
	const op = "repository.bucket.Delete"

	err := os.Remove(s.filePath + bucketName)

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

func (s *Store) IsEmptyBucket(bucketName string) (bool, error) {
	const op = "repository.bucket.IsEmpty"

	//pkg.ColsEqualValue(s.filePath + bucketName + "/" + "objects.csv")

	return true, nil
}
