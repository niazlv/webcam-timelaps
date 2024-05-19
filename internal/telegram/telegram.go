package telegram

type Messenger interface {
	SendVideo(videoPath string) error
}
