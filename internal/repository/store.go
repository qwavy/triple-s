package repository

import "sync"

type Store struct {
	mu                  sync.Mutex
	filePath            string
	bucketsInfoFilePath string
}

func NewStore(filePath, bucketsInfoFilePath string) *Store {
	return &Store{filePath: filePath, bucketsInfoFilePath: bucketsInfoFilePath}
}

func (s *Store) GetFilePath() string {
	return s.filePath
}
