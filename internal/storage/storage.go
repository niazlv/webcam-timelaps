package storage

// Interface
type Storage interface {
	SaveFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
	IsExist(path string) bool
	CreateDir(path string) error
	Remove(path string) error
	ListFiles(dir string) ([]string, error)
}
