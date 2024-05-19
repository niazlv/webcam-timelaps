package main

import (
	"fmt"
	"log"
	"niazlv/time-lapse/internal/camera"
	"niazlv/time-lapse/internal/config"
	"niazlv/time-lapse/internal/processing"
	"niazlv/time-lapse/internal/storage"
	"niazlv/time-lapse/internal/telegram"
	"path/filepath"
	"time"
)

func main() {
	var cam camera.Camera
	var store storage.Storage
	var msg telegram.Messenger

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}

	delay := cfg.Delay
	fmt.Println(delay)

	// Пример использования IP камеры D-Link и локального хранилища
	cam = &camera.IPCam{IP: cfg.IP, Username: cfg.Username, Password: cfg.Password}
	store = &storage.LocalStorage{Directory: cfg.Storage.Directory}
	msg = &telegram.TelegramBot{BotToken: cfg.Telegram.BotToken, ChatID: cfg.Telegram.ChatID, Storage: store}
	proc := &processing.Processing{Storage: store, OutputDir: cfg.VideoOutputDir}

	log.Println("Starting the main loop")
	go func() {
		for {
			err := captureAndSaveImage(cam, store)
			if err != nil {
				log.Printf("Error capturing and saving image: %v\n", err)
			}
			time.Sleep(delay)
		}
	}()

	// // Захват изображения
	// img, err := cam.CaptureImage()
	// if err != nil {
	// 	fmt.Println("Error capturing image:", err)
	// 	return
	// }

	// // Сохранение изображения
	// err = store.SaveFile("images/image.jpg", img)
	// if err != nil {
	// 	fmt.Println("Error saving image:", err)
	// 	return
	// }

	_ = msg
	_ = proc
	// // Отправка видео в Telegram
	// err = msg.SendVideo("images/image.jpg")
	// if err != nil {
	// 	fmt.Println("Error sending video:", err)
	// 	return
	// }

	// Планировщик задач для создания видео
	// c := cron.New()
	// _, err := c.AddFunc("@daily", func() {
	// 	now := time.Now()
	// 	todayDIR := now.Format("01-02-2006")
	// 	outputFile := filepath.Join(videoOutputDir, fmt.Sprintf("timelapse_%s.mp4", todayDIR))

	// 	log.Printf("Processing images to video: %s\n", outputFile)
	// 	err := proc.ProcessImagesToVideo(todayDIR, outputFile)
	// 	if err != nil {
	// 		log.Printf("Error processing video: %v\n", err)
	// 	} else {
	// 		log.Printf("Video created: %s\n", outputFile)
	// 		// Отправка видео в Telegram
	// 		// videoData, err := store.ReadFile(outputFile)
	// 		// if err != nil {
	// 		//     log.Printf("Error reading video file: %v\n", err)
	// 		// } else {
	// 		//     err = msg.SendVideo(videoData)
	// 		//     if err != nil {
	// 		//         log.Printf("Error sending video: %v\n", err)
	// 		//     }
	// 		// }
	// 	}
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to schedule video processing: %v\n", err)
	// }

	// c.Start()

	// go func() {
	// 	now := time.Now()
	// 	todayDIR := now.Format("01-02-2006")
	// 	outputFile := filepath.Join(videoOutputDir, fmt.Sprintf("timelapse_%s.mp4", todayDIR))

	// 	log.Printf("Processing images to video: %s\n", outputFile)
	// 	err := proc.ProcessImagesToVideo(todayDIR, outputFile)
	// 	if err != nil {
	// 		log.Printf("Error processing video: %v\n", err)
	// 	} else {
	// 		log.Printf("Video created: %s\n", outputFile)
	// 		// Отправка видео в Telegram
	// 		// videoData, err := store.ReadFile(outputFile)
	// 		// if err != nil {
	// 		//     log.Printf("Error reading video file: %v\n", err)
	// 		// } else {
	// 		//     err = msg.SendVideo(videoData)
	// 		//     if err != nil {
	// 		//         log.Printf("Error sending video: %v\n", err)
	// 		//     }
	// 		// }
	// 	}
	// }()

	// Блокировка выполнения программы
	select {}

	fmt.Println("Process completed successfully")
}

// Функция захвата и сохранения изображения
func captureAndSaveImage(cam camera.Camera, store storage.Storage) error {
	now := time.Now()
	todayDIR := now.Format("01-02-2006")
	todays := now.Format("15-04-05")
	filename := fmt.Sprintf("cam1_%s.jpg", todays)
	dirPath := todayDIR
	filePath := filepath.Join(dirPath, filename)

	// Проверка существования директории, создание если не существует
	if !store.IsExist(dirPath) {
		err := store.CreateDir(dirPath)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		log.Printf("Directory %s created\n", dirPath)
	}

	// Захват изображения
	img, err := cam.CaptureImage()
	if err != nil {
		return fmt.Errorf("failed to capture image: %w", err)
	}

	// Сохранение изображения
	err = store.SaveFile(filePath, img)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	log.Printf("Image saved to %s\n", filePath)
	return nil
}
