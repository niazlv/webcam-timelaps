package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Cam            CamConfig
	Delay          time.Duration
	Storage        StorageConfig
	Telegram       TelegramConfig
	VideoOutputDir string `mapstructure:"video_output_dir"`
}

type CamConfig struct {
	Type     string
	IP       string
	Username string
	Password string
	UsbCam   UsbCamConfig
	Axis     AxisConfig
}

type UsbCamConfig struct {
	Resolution string
	UsbDevice  string
}

type AxisConfig struct {
	Resolution  string
	Compression int
}

type StorageConfig struct {
	Type      string
	Directory string
	FTP       FTPConfig
}

type FTPConfig struct {
	Server   string
	Username string
	Password string
}

type TelegramConfig struct {
	BotToken string `mapstructure:"bot_token"`
	ChatID   string `mapstructure:"chat_id"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Настройка Viper
	v.SetConfigName("config") // Имя файла конфигурации без расширения
	v.SetConfigType("yaml")   // Тип файла конфигурации
	v.AddConfigPath("config") // Путь к конфигурационному файлу

	// Чтение основного конфигурационного файла
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Проверка наличия пользовательского конфигурационного файла
	v.SetConfigName("my_config") // Имя пользовательского файла конфигурации
	if err := v.MergeInConfig(); err != nil {
		log.Println("No my_config.yaml found, using default config.yaml")
	}

	// Считывание конфигурации в структуру
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Преобразование задержки из секунд в Duration
	config.Delay = config.Delay * time.Second

	return &config, nil
}
