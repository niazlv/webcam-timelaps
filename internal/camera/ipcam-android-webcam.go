package camera

import (
	"fmt"
	"os"
	"os/exec"
)

type IPCamAndroidWebCam struct {
	IP string
}

func (c *IPCamAndroidWebCam) CaptureImage() ([]byte, error) {
	//http://192.168.99.151:8080/photo.jpg
	url := fmt.Sprintf("http://%s/photo.jpg", c.IP)
	tmpFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	cmd := exec.Command("wget", url, "--timeout", "10", "-q", "-O", tmpFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture image from IP camera: %w", err)
	}

	imgData, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	return imgData, nil
}
