# Time-Lapse Project

This is a home project to create time-lapse videos from various types of cameras, including IP cameras and USB cameras. The captured images are processed into videos using `ffmpeg`, and the resulting videos are sent to a specified Telegram chat.

## Warning

Project in development, not for production!!!

## Features

- Capture images from IP cameras (Axis, D-Link) and USB cameras.
- Save images to local storage or an FTP server.
- Create directories and manage files in the storage.
- Process images into videos using `ffmpeg`.
- Send processed videos to a Telegram chat.
- Configuration using `config.yaml` and `my_config.yaml`.

## Requirements

- Go 1.16+
- `fswebcam` (for capturing images from USB cameras on Linux)
- `ffmpeg` (for processing images into videos)
- Access to an IP camera or USB camera
- Telegram bot token and chat ID

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-username/time-lapse-project.git
    cd time-lapse-project
    ```

2. Install required dependencies:

    ```sh
    sudo apt-get update
    sudo apt-get install fswebcam ffmpeg
    ```

3. Build the project:

    ```sh
    go build -o timelapse cmd/time-lapse/main.go
    ```

## Configuration

1. Copy the `config.yaml` to `my_config.yaml`:

    ```sh
    cp config.yaml my_config.yaml
    ```

2. Edit `my_config.yaml` with your configuration details.

## Usage

Run the application:

```sh
./timelapse
```

The application will start capturing images based on the configuration and process them into a video daily, sending the result to the specified Telegram chat.

## TODO List

- [ ] Add support for additional camera types.
- [ ] Improve error handling and logging.
- [ ] Implement unit tests for all components.
- [ ] Optimize the video processing pipeline.
- [ ] Add support for cloud storage services.
- [ ] Fix Usb camera support(usbcamera module not work corectly)
- [ ] Add webdav support
- [ ] Add smb support
- [ ] fix Processing support with Storage interface
  - because i'm using ffmpeg for make video from images, we need download images localy, for to do it. That's the problem. That there may not always be room or it may not always be possible. So we need to come up with a solution.
- [ ] Sending videos to a Telegram chat.
- [ ] Dockerize this app

## Working Components

- Image capture from IP cameras (Axis, D-Link, [android-ipcam](https://play.google.com/store/apps/details?id=com.pas.webcam)) and USB cameras(manual mode).
- Local and FTP storage support.
- Directory creation and file management in storage.
- Image processing into videos using `ffmpeg`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## Contact

For any questions or suggestions, please open an issue in the repository.

---

Thank you for using the Time-Lapse Project! Enjoy creating your time-lapse videos.
