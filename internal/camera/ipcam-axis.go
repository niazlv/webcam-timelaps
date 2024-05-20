package camera

import (
	"fmt"
	"os"
	"os/exec"
)

type IPCamAXIS struct {
	IP          string
	Username    string
	Password    string
	Resolution  string
	Compression int
}

func (c *IPCamAXIS) CaptureImage() ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s@%s/axis-cgi/jpg/image.cgi?%s&compression=%d", c.Username, c.Password, c.IP, c.Resolution, c.Compression)
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
