package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// LocalStorage реализует интерфейс Storage для локального хранилища
type LocalStorage struct {
	Directory string
}

// SaveFile сохраняет файл в локальное хранилище
func (s *LocalStorage) SaveFile(path string, data []byte) error {
	fullPath := filepath.Join(s.Directory, path)

	// Создаем директорию, если она не существует
	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return err
	}

	// Записываем данные в файл
	return os.WriteFile(fullPath, data, 0644)
}

func (s *LocalStorage) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.Directory, path))
}

func (s *LocalStorage) IsExist(path string) bool {
	fullpath := filepath.Join(s.Directory, path)
	_, err := os.Stat(fullpath)
	return !os.IsNotExist(err)
}

func (s *LocalStorage) CreateDir(path string) error {
	fullpath := filepath.Join(s.Directory, path)
	return os.Mkdir(fullpath, os.ModePerm)
}
func (s *LocalStorage) Remove(path string) error {
	return os.Remove(path)
}

func (s *LocalStorage) ListFiles(dir string) ([]string, error) {
	fullDir := filepath.Join(s.Directory, dir)
	files, err := os.ReadDir(fullDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, filepath.Join(fullDir, file.Name()))
		}
	}

	return fileNames, nil
}
