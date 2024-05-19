package camera

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

type USBCam struct {
	Device     string
	Resolution string
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "usb_cam_*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tempFile.Close()

	// Использование fswebcam для захвата изображения с USB камеры
	cmd := exec.Command("fswebcam", "-d", c.Device, "-r", c.Resolution, tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to capture image: %w\nOutput: %s", err, string(output))
	}

	var imgData []byte
	for {
		fileInfo, err := tempFile.Stat()
		if err != nil {
			return nil, fmt.Errorf("failed to get file info: %w", err)
		}

		if fileInfo.Size() > 0 {
			imgData, err = ioutil.ReadFile(tempFile.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to read captured image: %w", err)
			}
			break
		}
		time.Sleep(100 * time.Millisecond) // Подождать немного и проверить снова
	}

	return imgData, nil
}
