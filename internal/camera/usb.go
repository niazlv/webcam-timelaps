package camera

import "fmt"

type USBCam struct {
	Device string
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	// TODO: Реализация захвата изображения с USB камеры
	fmt.Println("Capturing image from USB camera")
	return []byte{}, nil
}
