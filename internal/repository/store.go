package repository

type Store struct {
	filePath            string
	bucketsInfoFilePath string
}

func NewStore(filePath, bucketsInfoFilePath string) *Store {
	return &Store{filePath: filePath, bucketsInfoFilePath: bucketsInfoFilePath}
}

func (s *Store) GetFilePath() string {
	return s.filePath
}
