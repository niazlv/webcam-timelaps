package storage

import "fmt"

type SMBStorage struct {
	URL      string
	Username string
	Password string
}

func (s *SMBStorage) SaveFile(path string, data []byte) error {
	// TODO: Реализация сохранения файла на SMB сервер
	fmt.Println("Saving file to SMB storage")
	return nil
}
