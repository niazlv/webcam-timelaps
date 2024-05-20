package camera

// Interface
type Camera interface {
	CaptureImage() ([]byte, error)
}
