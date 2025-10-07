# Movie API Test Documentation

## Tổng quan

Tài liệu này giải thích về file `test/movies_test.go`, cách viết test cho API Movie, các yêu cầu, cách mở rộng và tài liệu tham khảo.

## Cấu trúc file test

- **Imports**: Sử dụng các package chuẩn của Go, httptest, và các package của dự án (config, controller, cache, db).
- **Thiết lập test**:
  - Đổi thư mục làm việc về project root để load config.
  - Khởi tạo config, kết nối database, khởi tạo cache.
  - Tạo HTTP server với Gin và router của dự án.
- **Các case test**:
  - **POST /movies**: Tạo movie mới, kiểm tra response trả về đúng dữ liệu.
  - **GET /movies/{id}**: Lấy movie vừa tạo theo id.
  - **GET /movies/search**: Tìm kiếm movie theo title.

## Cách viết và chạy test

1. **Yêu cầu**:
   - Đã cài Go (`go1.18+`).
   - Đã chạy database và cache đúng config.
   - Đã cài [httpexpect](https://github.com/gavv/httpexpect).

2. **Chạy test**:
   ```sh
   go test -v ./test
   ```

3. **Cách viết test**:
   - Dùng `httptest.NewServer` để tạo server test.
   - Dùng `httpexpect` để gửi request và kiểm tra response.
   - Đóng server bằng `defer server.Close()`.

## Giải thích code

- `os.Chdir("..")`: Đảm bảo chạy test ở đúng thư mục để load config.
- `config.InitConfig()`: Load cấu hình (DB, cache, ...).
- `db.ConnectDB(cfg)`: Kết nối database.
- `cache.InitCache(cfg.MemcachedAddr)`: Khởi tạo cache.
- `controller.SetupRouter()`: Tạo router cho API.
- `httpexpect.Default(t, server.URL)`: Tạo client test.
- Các case test dùng `e.POST`, `e.GET` để gọi API và kiểm tra kết quả trả về.

## Mở rộng test

- Thêm test cho update, delete movie.
- Test các trường hợp lỗi (input sai, không tìm thấy, ...).
- Dùng fixture hoặc mock data cho các case phức tạp.
- Tích hợp CI/CD để tự động chạy test.

## Tham khảo

- [httpexpect GitHub](https://github.com/gavv/httpexpect)
- [Go testing documentation](https://pkg.go.dev/testing)
- [Gin web framework](https://github.com/gin-gonic/gin)
- [Go httptest package](https://pkg.go.dev/net/http/httptest)

