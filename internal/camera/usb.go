package camera

import (
	"fmt"
	"os"
	"os/exec"
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

	// Использование fswebcam для захвата изображения с USB камеры
	cmd := exec.Command("fswebcam", "-d", c.Device, "-r", c.Resolution, tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to capture image: %w\nOutput: %s", err, string(output))
	}

	// Чтение содержимого файла после завершения команды
	imgData, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read captured image: %w", err)
	}

	return imgData, nil
}
