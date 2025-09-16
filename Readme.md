# Tutorial Backend

## Kiến trúc 3 lớp (3-Layer Architecture)

### Controller (Presentation)
- Nhận request, xử lý nghiệp vụ và trả về response
- Chịu trách nhiệm cho việc xác thực và phân quyền
- Chuyển đổi dữ liệu giữa DTO (Data Transfer Objects) và model

### Model
- Chứa data model của service
- Định nghĩa các entity và mối quan hệ giữa chúng
- Chứa các method để validate dữ liệu

### Service
- Tương tác với cơ sở dữ liệu (ví dụ: PostgreSQL, MongoDB, Redis)
- Tích hợp với dịch vụ bên thứ ba (Twilio, email,...)
- Gọi đến các service khác trong hệ thống microservice

## Cấu trúc thư mục của HTTP API

```
/config
    config.go           # Cấu hình và biến môi trường
/controller
    main.go            # Định nghĩa router
    [resource].go      # Controller cho mỗi resource
    dto.go             # Data Transfer Objects
/model
    [resource].go      # Model cho mỗi resource
/service
    db/                # Tương tác với database
        [resource].go  # Repository cho mỗi resource
        db.go          # Khởi tạo kết nối DB
    [external].go      # Tích hợp với dịch vụ bên ngoài
main.go               # Entry point của ứng dụng
```

## Nguyên tắc thiết kế

1. **Dependency Rule**: Các lớp chỉ nên phụ thuộc vào lớp bên dưới hoặc các thư viện bên ngoài, không nên phụ thuộc vào lớp cùng cấp hoặc lớp trên

2. **Data Flow**: Luồng dữ liệu nên đi từ Controller → Service → Model và ngược lại

3. **DTO Usage**: DTO (Data Transfer Objects) chỉ nên được sử dụng ở lớp Controller, không nên truyền DTO xuống lớp Service hoặc Model

## Hướng dẫn sử dụng golang-migrate để tạo và chạy migration

### 1. Cài đặt golang-migrate

Bạn có thể cài đặt golang-migrate CLI bằng các cách sau:

```sh
# Cài đặt bằng Go (yêu cầu Go đã cài đặt)
go install -v github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Trên Windows (sử dụng scoop)
scoop install migrate

# Trên MacOS (sử dụng brew)
brew install migrate

# Hoặc tải trực tiếp từ https://github.com/golang-migrate/migrate/releases
```

Sau khi cài bằng go install, file thực thi migrate sẽ nằm ở thư mục $GOPATH/bin hoặc $HOME/go/bin. Bạn cần thêm thư mục này vào PATH để sử dụng lệnh migrate ở mọi nơi.

### 2. Tạo file migrate mới

Sử dụng lệnh sau để tạo file migrate:

```sh
migrate create -ext sql -dir ./migrations -seq ten_migration
```

- `-ext sql`: Định dạng file là .sql
- `-dir ./migrations`: Thư mục chứa các file migrate
- `-seq`: Tạo file với số thứ tự tăng dần
- `ten_migration`: Tên migration (ví dụ: create_users_table)

Sau khi chạy lệnh, bạn sẽ có 2 file:
- `xxxxxx_ten_migration.up.sql`: Chứa câu lệnh để migrate lên
- `xxxxxx_ten_migration.down.sql`: Chứa câu lệnh để rollback

### 3. Chạy migrate

Nếu bạn đã export/set biến môi trường DATABASE_URL thủ công trong shell, bạn có thể dùng lệnh migrate như sau:

```sh
$env:DB_URL="postgres://user:pass@localhost:5432/tutorial-go?sslmode=disable"
migrate -path ./service/db/migrations -database $env:DB_URL up

```

> **Lưu ý:** migrate CLI không tự động đọc file .env, nên nếu muốn dùng biến môi trường từ file .env, bạn cần dùng dotenv CLI hoặc nạp biến vào shell trước.

Tham khảo thêm tại: https://github.com/golang-migrate/migrate
