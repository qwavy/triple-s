package pkg

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func WriteDataToCsv(data []any, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := make([]string, len(data))
	for i, v := range data {
		fmt.Println(v)
		row[i] = fmt.Sprint(v)
	}

	return writer.Write(row)
}

func WriteDataToFile(content []byte, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func OverWriteDataToFile(content []byte, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func ReadFile(filePath string) ([]byte, error) {
	result, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindRowCsvByName(name string, index int, filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if record[index] == name {
			return record, nil
		}
	}
	return nil, fmt.Errorf("Error: Didn`t find row by name in csv")
}

func ColsEqualValue(value string, index int, filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	for _, record := range records {
		if record[index] != value {
			return false, nil
		}
	}

	return true, nil
}
