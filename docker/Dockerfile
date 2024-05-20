# стадия 1. Сборка проекта(не будет включен в финальный образ)
FROM golang:1.21-alpine as builder

# Ставим какие-либо зависимости системы
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache git build-base musl-dev alpine-sdk cmake clang clang-dev make gcc g++ libc-dev linux-headers

RUN mkdir /tmp/opencv
WORKDIR /tmp/opencv
RUN wget -O opencv.zip https://github.com/opencv/opencv/archive/4.9.0.zip
RUN unzip opencv.zip
RUN wget -O opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/4.9.0.zip
RUN unzip opencv_contrib.zip
RUN mkdir /tmp/opencv/opencv-4.9.0/build

WORKDIR /tmp/opencv/opencv-4.9.0/build
RUN cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D OPENCV_EXTRA_MODULES_PATH=/tmp/opencv/opencv_contrib-4.9.0/modules -D BUILD_DOCS=OFF BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=OFF -D BUILD_opencv_python=OFF -D BUILD_opencv_python2=OFF -D BUILD_opencv_python3=OFF ..
RUN make -j4
RUN make install

RUN rm -rf /tmp/opencv

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем модули Go(отдельно от исходников, чтобы лучше кешировалось и каждый раз по новой не обновлялось)
COPY go.mod .
COPY go.sum .

# Скачиваем зависимости уже GO
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение с поддержкой CGO
ENV CGO_ENABLED=1
RUN go build -o main ./cmd/time-lapse/main.go


# стадия 2. Запуск в отдельном контейнере
FROM alpine:3.14

RUN mkdir /app

# костыль доступа к storage host машины
ARG UID=1000
ARG GID=1000

RUN adduser -D -u ${UID} -g ${GID} appuser && mkdir /app-data && chown -R appuser /app-data

RUN mkdir /app-data/storage && chown -R appuser /app-data/storage && ln -s /app-data/storage /app/storage

WORKDIR /app

COPY --from=builder --chown=root:root /build/main .

RUN chmod 755 main

USER appuser

# Указываем команду для запуска приложения
CMD ["/app/main"]