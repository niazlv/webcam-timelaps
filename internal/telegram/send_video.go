package telegram

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"niazlv/time-lapse/internal/storage"
	"path/filepath"
)

type TelegramBot struct {
	BotToken string
	ChatID   string
	Storage  storage.Storage
}

func (t *TelegramBot) SendVideo(videoPath string) error {
	// Читаем видео файл из хранилища
	fileData, err := t.Storage.ReadFile(videoPath)
	if err != nil {
		return fmt.Errorf("could not read video file: %w", err)
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Добавляем видео файл в форму
	part, err := writer.CreateFormFile("video", filepath.Base(videoPath))
	if err != nil {
		return fmt.Errorf("could not create form file: %w", err)
	}

	_, err = part.Write(fileData)
	if err != nil {
		return fmt.Errorf("could not write file data to form: %w", err)
	}

	// Добавляем остальные параметры формы
	writer.WriteField("chat_id", t.ChatID)

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("could not close writer: %w", err)
	}

	// Формируем URL для запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", t.BotToken)

	// Создаем HTTP запрос
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Video sent to Telegram")
	return nil
}
