package tests

import (
	"log"
	"niazlv/time-lapse/internal/config"
	"niazlv/time-lapse/internal/storage"
	"os"
	"path/filepath"
	"testing"
)

var cfg config.Config

// Загрузка конфигурации

func TestLocalStorage_SaveFile(t *testing.T) {
	// Создаем временную директорию для теста
	tempDir := t.TempDir()

	// Создаем экземпляр LocalStorage
	localStorage := &storage.LocalStorage{Directory: tempDir}

	// Данные для записи
	filePath := "test/testfile.txt"
	data := []byte("Hello, World!")

	// Пытаемся сохранить файл
	err := localStorage.SaveFile(filePath, data)
	if err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}

	// Проверяем, что файл был сохранен
	fullPath := tempDir + "/" + filePath
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		t.Fatalf("File does not exist: %v", err)
	}

	// Проверяем содержимое файла
	savedData, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}
	if string(savedData) != string(data) {
		t.Fatalf("Saved data does not match: got %v, want %v", string(savedData), string(data))
	}
}

func TestLocalStorage_ReadFile(t *testing.T) {
	// Создаем временную директорию для теста
	tempDir := t.TempDir()

	// Создаем экземпляр LocalStorage
	localStorage := &storage.LocalStorage{Directory: tempDir}

	// Данные для записи
	filePath := "test/testfile.txt"
	data := []byte("Hello, World!")

	// Пытаемся сохранить файл
	err := localStorage.SaveFile(filePath, data)
	if err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}

	// Проверяем, что файл был сохранен
	fullPath := tempDir + "/" + filePath
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		t.Fatalf("File does not exist: %v", err)
	}

	// Проверяем содержимое файла
	savedData, err := localStorage.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}
	if string(savedData) != string(data) {
		t.Fatalf("Saved data does not match: got %v, want %v", string(savedData), string(data))
	}
}

func TestLocalStorage_IsExist(t *testing.T) {
	// Создаем временную директорию для теста
	tempDir := t.TempDir()

	// Создаем экземпляр LocalStorage
	localStorage := &storage.LocalStorage{Directory: tempDir}

	// Данные для записи
	filePath := "test/testfile.txt"
	data := []byte("Hello, World!")

	// Пытаемся сохранить файл
	err := localStorage.SaveFile(filePath, data)
	if err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}

	// Проверяем, что файл был сохранен
	if !localStorage.IsExist(filePath) {
		t.Fatalf("File does not exist: %v", err)
	}

	// Проверяем содержимое файла
	savedData, err := localStorage.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}
	if string(savedData) != string(data) {
		t.Fatalf("Saved data does not match: got %v, want %v", string(savedData), string(data))
	}
}

func TestFTPStorage_SaveFileAndReadFile(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}
	// Создаем экземпляр LocalStorage
	ftpStorage := &storage.FTPStorage{Server: cfg.Storage.FTP.Server, Username: cfg.Storage.FTP.Username, Password: cfg.Storage.FTP.Password}

	// Данные для записи
	dirpath := "__test_temp_dir__"
	filePath := "testfile.txt"
	data := []byte("Hello, World!")

	err = ftpStorage.Remove(dirpath)
	if err != nil {
		t.Logf("good, storage not exist(maybe). Check it self: %v", err)
	}
	err = ftpStorage.CreateDir(dirpath)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	// Пытаемся сохранить файл
	fullpath := filepath.Join(dirpath, filePath)
	err = ftpStorage.SaveFile(fullpath, data)
	if err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}
	// Добавим небольшую задержку
	// time.Sleep(2 * time.Second)

	// Проверяем, что файл был сохранен
	if !ftpStorage.IsExist(fullpath) {
		t.Fatalf("File does not exist: %v", err)
	}

	// Проверяем содержимое файла
	savedData, err := ftpStorage.ReadFile(fullpath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}
	if string(savedData) != string(data) {
		t.Fatalf("Saved data does not match: got %v, want %v", string(savedData), string(data))
	}

	err = ftpStorage.Remove(dirpath)
	if err != nil {
		t.Fatalf("Failed to delete saved file: %v", err)
	}
}
