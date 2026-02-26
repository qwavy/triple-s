package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
)

func (s *Store) IsExistsBucket(name string) (bool, error) {
	const op = "repository.bucket.Exists"

	filePathCsv := filepath.Join(s.filePath, "buckets.csv")

	file, err := os.Open(filePathCsv)
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
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.bucket.Save"

	filePathCsv := filepath.Join(s.filePath, "buckets.csv")
	bucketPath := s.filePath + b.Name

	err := os.Mkdir(bucketPath, 0755)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	creationTime := pkg.GetTime()
	err = pkg.WriteDataToCsv([]any{b.Name, creationTime, "active"}, filePathCsv)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Store) ListBucket() ([]models.Bucket, error) {
	const op = "repository.bucket.List"

	filePathCsv := filepath.Join(s.filePath, "buckets.csv")

	file, err := os.Open(filePathCsv)
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
	s.mu.Lock()
	defer s.mu.Unlock()

	const op = "repository.bucket.Delete"

	filePathCsv := filepath.Join(s.filePath, "buckets.csv")

	err := os.Remove(filePathCsv)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	file, err := os.Open(filePathCsv)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	file, err = os.Create(filePathCsv + ".tmp")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, record := range records {
		if record[0] != bucketName {
			err := writer.Write(record)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		}

	}
	writer.Flush()

	return os.Rename(filePathCsv+".tmp", filePathCsv)
}

func (s *Store) IsEmptyBucket(bucketName string) (bool, error) {
	const op = "repository.bucket.IsEmpty"

	//pkg.ColsEqualValue(s.filePath + bucketName + "/" + "objects.csv")

	return true, nil
}

func (s *Store) LastModificationTime(bucketName string) error {
	const op = "repository.object.LastModificationTime"

	filePathCSV := "./data/buckets.csv"

	file, err := os.Open(filePathCSV)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	file, err = os.Create(filePathCSV + ".tmp")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, record := range records {
		if record[0] != bucketName {
			err := writer.Write(record)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		} else {
			creationTime := pkg.GetTime()
			newRecord := []string{record[0], record[1], creationTime, "true"}
			err := writer.Write(newRecord)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		}

	}
	writer.Flush()

	return os.Rename(filePathCSV+".tmp", filePathCSV)
}
