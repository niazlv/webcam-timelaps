package camera

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type USBCam struct {
	Device     string
	Resolution string
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	tempFile, err := os.CreateTemp("", "usb_cam_*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Удаление временного файла после использования
	fmt.Println(tempFile.Name())

	// Использование fswebcam для захвата изображения с USB камеры
	cmd := exec.Command("fswebcam", "-d", c.Device, "-r", c.Resolution, tempFile.Name())
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start fswebcam: %w", err)
	}

	// Установить таймаут на ожидание завершения команды
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(3 * time.Second):
		// Таймаут
		if err := cmd.Process.Kill(); err != nil {
			return nil, fmt.Errorf("failed to kill process after timeout: %w", err)
		}
		return nil, fmt.Errorf("timed out while waiting for fswebcam to finish")
	case err := <-done:
		// Команда завершилась
		if err != nil {
			output, _ := cmd.CombinedOutput()
			return nil, fmt.Errorf("fswebcam failed: %w\nOutput: %s", err, string(output))
		}
	}

	// Проверка, что файл действительно заполнен
	var imgData []byte
	start := time.Now()
	for {
		fileInfo, err := os.Stat(tempFile.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to get file info: %w", err)
		}

		if fileInfo.Size() > 0 {
			imgData, err = os.ReadFile(tempFile.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to read captured image: %w", err)
			}
			break
		}

		if time.Since(start) > 3*time.Second {
			return nil, fmt.Errorf("timed out while waiting for image file to be created")
		}

		time.Sleep(100 * time.Millisecond) // Подождать немного и проверить снова
	}

	return imgData, nil
}
