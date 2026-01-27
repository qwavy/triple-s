package storage

type ObjectStorage struct {
	filePath string
}

func newObjectStorage(path string) *ObjectStorage{
	return &{ObjectStorage{filePath: path}}
}