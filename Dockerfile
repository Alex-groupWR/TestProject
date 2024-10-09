# Используем официальный образ Go как базовый
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY . .
# Скачиваем зависимости
RUN go mod tidy

# Копируем остальные файлы проекта
COPY go.mod go.sum main.go ./



# Собираем приложение
RUN go build -o main

# Указываем команду для запуска приложения
CMD ["./main"]
