package camera

import (
	"fmt"
	"os"
	"os/exec"
)

type IPCamDLink struct {
	IP       string
	Username string
	Password string
}

func (c *IPCamDLink) CaptureImage() ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s@%s/dms?nowprofileid=1", c.Username, c.Password, c.IP)

	tmpFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	cmd := exec.Command("wget", url, "-q", "-O", tmpFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture image from IP camera: %w", err)
	}

	imgData, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	return imgData, nil
}
