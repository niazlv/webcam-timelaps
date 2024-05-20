package processing

import (
	"fmt"
	"niazlv/time-lapse/internal/storage"
	"os/exec"
	"path/filepath"
)

type Processing struct {
	Storage   storage.Storage
	OutputDir string
}

func (p *Processing) ProcessImagesToVideo(imageDir, outputFile string) error {
	// Проверка существования директории
	if !p.Storage.IsExist(imageDir) {
		return fmt.Errorf("directory does not exist: %s", imageDir)
	}

	imageFiles, err := p.Storage.ListFiles(imageDir)
	if err != nil {
		return fmt.Errorf("failed to list image files: %w", err)
	}

	if len(imageFiles) == 0 {
		return fmt.Errorf("no images found in directory: %s", imageDir)
	}

	// Проверка наличия хотя бы одного файла, соответствующего шаблону cam1_*.jpg
	matched, err := filepath.Glob(filepath.Join(imageDir, "cam1_*.jpg"))
	if err != nil || len(matched) == 0 {
		return fmt.Errorf("no images matching pattern found in directory: %s", imageDir)
	}

	// Запуск ffmpeg для создания видео
	cmd := exec.Command("ffmpeg", "-pattern_type", "glob", "-i", filepath.Join(imageDir, "cam1_*.jpg"), "-vf", "fps=25,format=yuv420p", outputFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create video: %w", err)
	}

	return nil
}
