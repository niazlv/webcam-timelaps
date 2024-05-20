package storage

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"io"
	"path/filepath"

	"github.com/jlaffaye/ftp"
)

type FTPStorage struct {
	Server   string
	Username string
	Password string
}

func (s *FTPStorage) connect() (*ftp.ServerConn, error) {
	conn, err := ftp.Dial(s.Server, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	err = conn.Login(s.Username, s.Password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *FTPStorage) SaveFile(path string, data []byte) error {
	conn, err := s.connect()
	if err != nil {
		return fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	reader := bytes.NewReader(data)

	// Создаем директорию, если она не существует
	dir := filepath.Dir(path)
	if err := s.createDirectories(conn, dir); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	err = conn.Stor(path, reader)
	if err != nil {
		return fmt.Errorf("failed to upload file to FTP server: %w", err)
	}

	return nil
}

func (s *FTPStorage) ReadFile(path string) ([]byte, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	response, err := conn.Retr(path)
	if err != nil {
		return nil, fmt.Errorf("failed to download file from FTP server: %w", err)
	}
	defer response.Close()

	data, err := io.ReadAll(response)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	return data, nil
}

func (s *FTPStorage) IsExist(path string) bool {
	conn, err := s.connect()
	if err != nil {
		fmt.Printf("failed to connect to FTP server: %v\n", err)
		return false
	}
	defer conn.Quit()

	// Try to retrieve the file
	_, err = conn.Retr(path)
	if err == nil {
		// File exists
		return true
	}

	// If retrieving the file fails, check if it's a directory
	entries, err := conn.List(path)
	if err == nil && len(entries) > 0 {
		// Directory exists
		return true
	}

	return false
}

func (s *FTPStorage) CreateDir(path string) error {
	conn, err := s.connect()
	if err != nil {
		return fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	return conn.MakeDir(path)
}

func (s *FTPStorage) Remove(path string) error {
	conn, err := s.connect()
	if err != nil {
		return fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	return s.removeRecursive(conn, path)
}

func (s *FTPStorage) removeRecursive(conn *ftp.ServerConn, path string) error {
	// Try to delete file
	err := conn.Delete(path)
	if err == nil {
		return nil
	}

	// List contents of the directory
	entries, err := conn.List(path)
	if err != nil {
		return fmt.Errorf("failed to list directory contents: %w", err)
	}

	for _, entry := range entries {
		// Avoid infinite loops caused by '.' and '..'
		if entry.Name == "." || entry.Name == ".." {
			continue
		}

		entryPath := fmt.Sprintf("%s/%s", path, entry.Name)
		if entry.Type == ftp.EntryTypeFile {
			err = conn.Delete(entryPath)
		} else if entry.Type == ftp.EntryTypeFolder {
			err = s.removeRecursive(conn, entryPath)
		}

		if err != nil {
			return fmt.Errorf("failed to delete %s: %w", entryPath, err)
		}
	}

	// Now remove the directory itself
	err = conn.RemoveDir(path)
	if err != nil {
		return fmt.Errorf("failed to delete directory %s: %w", path, err)
	}
	return nil
}

func (s *FTPStorage) createDirectories(conn *ftp.ServerConn, path string) error {
	parts := strings.Split(path, "/")
	for i := range parts {
		dir := strings.Join(parts[:i+1], "/")
		if dir == "" {
			continue
		}
		if err := conn.MakeDir(dir); err != nil && !ftpErrAlreadyExists(err) {
			return err
		}
	}
	return nil
}

func ftpErrAlreadyExists(err error) bool {
	// FTP server returns 550 error code if the directory already exists
	return err != nil && strings.Contains(err.Error(), "550")
}

func (s *FTPStorage) ListFiles(dirPath string) ([]string, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	entries, err := conn.List(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFile {
			files = append(files, filepath.Join(dirPath, entry.Name))
		}
	}

	return files, nil
}
