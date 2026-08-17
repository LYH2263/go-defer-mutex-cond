# 评测专用：保留完整 Go 工具链
FROM golang:1.22
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
CMD ["go", "test", "./...", "-count=1"]
