package tests

import (
	"niazlv/time-lapse/internal/storage"
	"os"
	"path/filepath"
	"testing"
)

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
	// Создаем экземпляр LocalStorage
	ftpStorage := &storage.FTPStorage{Server: "YOUR_FTP_SERVER", Username: "YOUR_LOGIN", Password: "YOUR_PASS"}

	// Данные для записи
	dirpath := "__test_temp_dir__"
	filePath := "testfile.txt"
	data := []byte("Hello, World!")

	err := ftpStorage.Remove(dirpath)
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
