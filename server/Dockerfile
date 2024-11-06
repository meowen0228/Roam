# 使用官方 Go 映像作為建置階段
FROM golang:1.23-alpine AS builder

# 設置工作目錄
WORKDIR /app

# 將Go模組文件複製到容器中
COPY go.mod .
COPY go.sum .

# 下載所有依賴項
RUN go mod download

# 將源代碼複製到容器中
COPY . .

# 構建應用
RUN go build -o main .

# 啟動應用
CMD ["./main"]
