# Hướng dẫn chạy ứng dụng

## 1. Cài đặt các phụ thuộc

```sh
go mod tidy
```

## 2. Thiết lập biến môi trường

- Tạo file `.env` hoặc export biến môi trường cần thiết (xem ví dụ trong `config/config.go` hoặc file `.env.example` nếu có).

## 3. Chạy database (Postgres, Memcached...)

- Có thể dùng Docker Compose hoặc tự chạy từng service:

```sh
docker-compose up -d
```

- Hoặc chạy từng container:

```sh
docker run --name tutorial-postgres -e POSTGRES_PASSWORD=yourpassword -e POSTGRES_DB=tutorial -p 5432:5432 -d postgres:15
```

## 4. Chạy migrate database

- Xem hướng dẫn chi tiết trong `docs/migrate.md`.

## 5. Chạy ứng dụng

```sh
go run main.go
```

## 6. Truy cập API

- Mặc định ứng dụng chạy ở port 3000 hoặc 8080 (tùy cấu hình). Truy cập các endpoint như:
  - `GET /healthz`
  - `GET /movies/:id`
  - `POST /movies`
  - ...

## 7. Chạy test

```sh
go test -v ./test
```

## 8. Build Docker image (nếu cần)

```sh
docker build -t tutorial-backend:latest .
```

## 9. Triển khai với Kubernetes

- Xem hướng dẫn chi tiết trong `docs/k8s.md`.

---

## Tham khảo
- Kiến trúc: `docs/architecture.md`
- Migration: `docs/migrate.md`
- Caching: `docs/Caching_Memcached.md`
- Logging & Metrics: `docs/Logging_Metrics.md`
- Paging: `docs/Paging.md`
- Test API: `docs/test_movies_api.md`