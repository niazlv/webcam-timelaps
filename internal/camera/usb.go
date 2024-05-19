package camera

import (
	"fmt"

	"gocv.io/x/gocv"
)

type USBCam struct {
	Device int
}

func (c *USBCam) CaptureImage() ([]byte, error) {
	// Открыть устройство камеры
	webcam, err := gocv.OpenVideoCapture(c.Device)
	if err != nil {
		return nil, fmt.Errorf("error opening video capture device: %v", c.Device)
	}
	defer webcam.Close()

	// Создать матрицу для хранения изображения
	img := gocv.NewMat()
	defer img.Close()

	// Захватить изображение
	if ok := webcam.Read(&img); !ok {
		return nil, fmt.Errorf("cannot read from device %v", c.Device)
	}

	// Проверить, что изображение действительно захвачено
	if img.Empty() {
		return nil, fmt.Errorf("image is empty")
	}

	// Преобразовать изображение в формат []byte
	buf, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return nil, fmt.Errorf("error encoding image: %v", err)
	}

	return buf.GetBytes(), nil
}
