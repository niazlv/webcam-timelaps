package camera

import "fmt"

type ADBCam struct {
	Device string
}

func (c *ADBCam) CaptureImage() ([]byte, error) {
	// TODO: Реализация захвата изображения с камеры телефона через ADB
	fmt.Println("Capturing image from ADB camera")
	return []byte{}, nil
}
