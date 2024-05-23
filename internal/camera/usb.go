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
	fileName := "usb_cam_capture.jpg"

	// Удаляем файл, если он существует
	if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to remove existing file: %w", err)
	}

	// Явно указываем путь к fswebcam (проверьте, что fswebcam доступен по этому пути)
	cmd := exec.Command("/usr/bin/fswebcam", "-d", c.Device, "-r", c.Resolution, fileName)

	// Перенаправляем стандартный вывод и стандартный вывод ошибок
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run fswebcam: %w", err)
	}

	// Проверка, что файл действительно заполнен с таймаутом
	var imgData []byte
	start := time.Now()
	for {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			return nil, fmt.Errorf("failed to get file info: %w", err)
		}

		if fileInfo.Size() > 0 {
			imgData, err = os.ReadFile(fileName)
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
