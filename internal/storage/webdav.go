package storage

import "fmt"

type WebDAVStorage struct {
	URL      string
	Username string
	Password string
}

func (s *WebDAVStorage) SaveFile(path string, data []byte) error {
	// TODO: Реализация сохранения файла на WebDAV сервер
	fmt.Println("Saving file to WebDAV storage")
	return nil
}
