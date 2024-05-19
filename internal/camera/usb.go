package camera

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type USBCam struct {
	Device string
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "usb_cam_*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tempFile.Close()

	// Использование ffmpeg для захвата изображения с USB камеры
	cmd := exec.Command("ffmpeg", "-f", "video4linux2", "-i", c.Device, "-vframes", "1", tempFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture image: %w", err)
	}

	imgData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read captured image: %w", err)
	}

	return imgData, nil
}
