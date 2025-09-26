# Hướng dẫn dựng Dockerfile và chạy ứng dụng với Docker

## 1. Dockerfile là gì?
Dockerfile là file mô tả các bước để xây dựng một Docker image cho ứng dụng. File này không có đuôi mở rộng, tên chuẩn là `Dockerfile`.

## 2. Ví dụ về Dockerfile cho ứng dụng Go
```dockerfile
# Sử dụng image Go chính thức để build
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Tạo image chạy nhẹ từ binary đã build
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
```

### Giải thích các lệnh:
- `FROM golang:1.21 AS builder`: Dùng image Go để build ứng dụng.
- `WORKDIR /app`: Đặt thư mục làm việc là `/app`.
- `COPY . .`: Copy toàn bộ mã nguồn vào image.
- `RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go`: Build file Go thành binary tên `app`.
- `FROM alpine:latest`: Dùng image Alpine nhẹ để chạy app.
- `WORKDIR /root/`: Đặt thư mục làm việc là `/root/`.
- `COPY --from=builder /app/app .`: Copy binary từ image build sang image chạy.
- `EXPOSE 8080`: Mở cổng 8080 cho container.
- `CMD ["./app"]`: Lệnh chạy khi container khởi động.

## 3. Cách build và chạy Docker image

### Build image
```sh
docker build -t my-go-app .
```
- `-t my-go-app`: Đặt tên cho image là `my-go-app`.
- `.`: Build từ thư mục hiện tại (chứa Dockerfile).

### Chạy container
```sh
docker run -p 8080:8080 my-go-app
```
- `-p 8080:8080`: Map cổng 8080 của máy host sang container.
- `my-go-app`: Tên image vừa build.

## 3.1. Đặt tag cho Docker image

Tag giúp bạn phân biệt các phiên bản image (ví dụ: latest, v1.0.0, dev, prod...). Tag được đặt khi build image bằng tuỳ chọn `-t`.

### Ví dụ đặt tag khi build
```sh
docker build -t my-go-app:latest .
docker build -t my-go-app:v1.0.0 .
```
- `my-go-app:latest`: Tag mặc định, thường dùng cho phiên bản mới nhất.
- `my-go-app:v1.0.0`: Tag cho phiên bản cụ thể.

Bạn có thể dùng tag để quản lý nhiều phiên bản image, deploy cho các môi trường khác nhau.

## 4. Một số lệnh Docker phổ biến
- `docker ps`: Xem các container đang chạy.
- `docker images`: Xem các image đã build.
- `docker stop <container_id>`: Dừng container.
- `docker rm <container_id>`: Xóa container.

## 5. Lưu ý
- Dockerfile phải đặt tên là `Dockerfile`, không có đuôi `.yml`.
- Nên dùng multi-stage build để tạo image nhỏ, bảo mật hơn.
- Có thể dùng thêm docker-compose nếu muốn chạy nhiều service cùng lúc.

---
Tham khảo thêm: https://docs.docker.com/engine/reference/builder/
