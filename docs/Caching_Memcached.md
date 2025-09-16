# Bài 4: Caching với Memcached

## Mục tiêu
- Giảm tải DB bằng cache.

## Khái niệm về Cache
Cache là vùng lưu trữ tạm thời giúp truy xuất dữ liệu nhanh hơn bằng cách lưu lại kết quả truy vấn hoặc dữ liệu thường dùng. Khi ứng dụng cần dữ liệu, nó kiểm tra cache trước, nếu có thì lấy ra (cache hit), nếu không thì truy vấn DB (cache miss) và lưu lại vào cache cho lần sau.

### Lợi ích
- Giảm tải cho database
- Tăng tốc độ phản hồi
- Giảm chi phí tài nguyên

### Các loại cache
- In-memory cache: Lưu trong RAM (Memcached, Redis)
- Distributed cache: Chia sẻ giữa nhiều server
- Negative cache: Lưu kết quả không tồn tại

## Memcached
Memcached là hệ thống cache phân tán, lưu dữ liệu dạng key-value trong RAM, rất nhanh và phù hợp cho cache tạm thời.
- Trang chủ: https://memcached.org/
- Go client: github.com/bradfitz/gomemcache/memcache

## Các khái niệm liên quan
- TTL (Time To Live): Thời gian dữ liệu tồn tại trong cache
- Cache Invalidation: Xóa hoặc cập nhật cache khi dữ liệu gốc thay đổi
- Cache Hit: Lấy dữ liệu thành công từ cache
- Cache Miss: Không tìm thấy dữ liệu trong cache
- Negative Caching: Lưu lại kết quả không tồn tại để tránh truy vấn lại

## Quy trình caching cho API
- GET /movies/{id}:
  - Kiểm tra cache với key là movie id
  - Nếu có dữ liệu (cache hit): trả về và log
  - Nếu không có (cache miss): truy vấn DB, lưu vào cache TTL 5 phút
  - Nếu movie không tồn tại: lưu negative cache TTL 30s
- POST /movies:
  - Invalidate (xóa) cache liên quan đến movie vừa tạo

## Ví dụ sử dụng Memcached trong Golang

### Cách 1: Dùng trực tiếp client gomemcache

```go
import "tutorial/service/cache"

// Lấy dữ liệu từ cache
item, err := cache.Client.Get("movie_123")
if err != nil {
    // xử lý lỗi
}
data := item.Value

// Lưu dữ liệu vào cache
err = cache.Client.Set(&memcache.Item{Key: "movie_123", Value: []byte("data"), Expiration: 300})

// Xóa dữ liệu khỏi cache
err = cache.Client.Delete("movie_123")
```

### Cách 2: Dùng hàm tiện ích wrapper

```go
import "tutorial/service/cache"

// Lấy dữ liệu từ cache
data, err := cache.Get("movie_123")
if err != nil {
    // xử lý lỗi
}

// Lưu dữ liệu vào cache
err = cache.Set("movie_123", []byte("data"), 300)

// Xóa dữ liệu khỏi cache
err = cache.Delete("movie_123")
```

**So sánh:**
- Cách 2 giúp code ngắn gọn, dễ đọc, dễ mở rộng logic (thêm log, negative cache, ...).
- Cách 1 phải thao tác trực tiếp với client, code dài hơn và khó mở rộng.

Bạn nên dùng cách 2 để code sạch, dễ bảo trì và mở rộng cho các tính năng cache nâng cao.

## Tài liệu tham khảo
- https://memcached.org/about/index.html
- https://github.com/bradfitz/gomemcache
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Caching

## Từ khóa tra cứu
- golang memcached client
- cache invalidation
- negative caching
