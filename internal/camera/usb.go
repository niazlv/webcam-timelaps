package camera

import (
	"fmt"

	"github.com/blackjack/webcam"
)

type USBCam struct {
	Device string
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	cam, err := webcam.Open(c.Device)
	if err != nil {
		return nil, fmt.Errorf("failed to open webcam: %w", err)
	}
	defer cam.Close()

	formatDesc := cam.GetSupportedFormats()
	var format webcam.PixelFormat
	for f := range formatDesc {
		format = f
		break
	}

	width, height := uint32(640), uint32(480)
	if _, _, _, err = cam.SetImageFormat(format, width, height); err != nil {
		return nil, fmt.Errorf("failed to set image format: %w", err)
	}

	if err := cam.StartStreaming(); err != nil {
		return nil, fmt.Errorf("failed to start streaming: %w", err)
	}
	defer cam.StopStreaming()

	timeout := uint32(5)
	for {
		err = cam.WaitForFrame(timeout)
		switch err.(type) {
		case nil:
			frame, err := cam.ReadFrame()
			if len(frame) != 0 {
				return frame, nil
			}
			if err != nil {
				return nil, fmt.Errorf("failed to read frame: %w", err)
			}
		case *webcam.Timeout:
			continue
		default:
			return nil, fmt.Errorf("error while waiting for frame: %w", err)
		}
	}
}
