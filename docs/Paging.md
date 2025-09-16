# Paging: Offset vs Cursor

## Offset Paging

- **Khái niệm**: Offset paging là cách phân trang truyền thống, sử dụng chỉ số bắt đầu (offset) và số lượng bản ghi (limit) để truy vấn dữ liệu.
- **Cách hoạt động**: Ví dụ, truy vấn `SELECT * FROM movies LIMIT 10 OFFSET 20` sẽ lấy 10 bản ghi bắt đầu từ vị trí thứ 21.
- **Ưu điểm**:
  - Đơn giản, dễ hiểu và dễ triển khai.
  - Phù hợp với dữ liệu nhỏ hoặc không thay đổi thường xuyên.
- **Nhược điểm**:
  - Khi dữ liệu thay đổi (thêm/xóa), kết quả có thể bị lặp hoặc thiếu.
  - Hiệu năng kém với dữ liệu lớn vì phải duyệt qua nhiều b���n ghi để đến offset.

## Cursor Paging

- **Khái niệm**: Cursor paging sử dụng một con trỏ (cursor) đại diện cho vị trí hiện tại trong tập dữ liệu, thường là giá trị của một trường duy nhất (ví dụ: id, timestamp).
- **Cách hoạt động**: Truy vấn sẽ lấy các bản ghi sau giá trị cursor, ví dụ: `SELECT * FROM movies WHERE id > last_id LIMIT 10`.
- **Ưu điểm**:
  - Hiệu năng tốt với dữ liệu lớn, không cần duyệt qua toàn bộ bản ghi.
  - Kết quả ổn định hơn khi dữ liệu thay đổi.
- **Nhược điểm**:
  - Phức tạp hơn khi triển khai.
  - Không phù hợp với dữ liệu không có trường duy nhất hoặc không có thứ tự rõ ràng.

# So sánh chi tiết Offset Paging và Cursor Paging

### 1. Hiệu năng
- **Offset Paging**: Khi offset lớn, database phải duyệt qua nhiều bản ghi để đến vị trí mong muốn, gây chậm và tốn tài nguyên. Đặc biệt với bảng lớn, truy vấn sẽ càng chậm.
- **Cursor Paging**: Truy vấn chỉ cần biết giá trị cursor (ví dụ: id cuối cùng), database có thể sử dụng index để lấy bản ghi tiếp theo nhanh chóng, hiệu năng cao hơn nhiều.

### 2. Độ ổn định dữ liệu
- **Offset Paging**: Nếu dữ liệu bị thêm/xóa trong quá trình phân trang, kết quả có thể bị lặp hoặc thiếu bản ghi, dẫn đến trải nghiệm không nhất quán cho người dùng.
- **Cursor Paging**: Kết quả ổn định hơn, vì luôn lấy bản ghi sau cursor, tránh lặp/thiếu khi dữ liệu thay đổi.

### 3. Tính mở rộng
- **Offset Paging**: Không phù hợp với dữ liệu lớn hoặc hệ thống cần mở rộng, vì hiệu năng giảm mạnh khi offset tăng.
- **Cursor Paging**: Phù hợp với hệ thống lớn, dữ liệu nhiều, dễ mở rộng và tối ưu hóa với index.

### 4. Độ phức tạp triển khai
- **Offset Paging**: Dễ triển khai, chỉ cần truyền offset và limit.
- **Cursor Paging**: Cần xác định trường làm cursor (id, timestamp...), truyền giá trị cursor, phức tạp hơn khi xử lý các trường hợp đặc biệt (ví dụ: sắp xếp theo nhiều trường).

### 5. Trường hợp sử dụng thực tế
- **Offset Paging**: Phù hợp cho admin, báo cáo, hoặc dữ liệu nhỏ, ít thay đổi.
- **Cursor Paging**: Phù hợp cho API public, feed mạng xã hội, dữ liệu lớn, cần hiệu năng và độ ổn định cao.

### 6. Ví dụ minh họa
- **Offset Paging**:
  ```sql
  SELECT * FROM movies LIMIT 10 OFFSET 20;
  ```
  Lấy 10 bản ghi từ vị trí thứ 21.

- **Cursor Paging**:
  ```sql
  SELECT * FROM movies WHERE id > 20 LIMIT 10;
  ```
  Lấy 10 bản ghi có id lớn hơn 20.

---

## Bảng so sánh tổng hợp
| Tiêu chí         | Offset Paging         | Cursor Paging         |
|------------------|----------------------|----------------------|
| Đơn giản         | ✔️                   | ❌                   |
| Hiệu năng        | ❌                   | ✔️                   |
| Ổn định dữ liệu  | ❌                   | ✔️                   |
| Mở rộng          | ❌                   | ✔️                   |
| Triển khai       | Dễ                   | Khó hơn              |
| Trường hợp dùng  | Dữ liệu nhỏ          | Dữ liệu lớn          |

## Kết luận
- Offset paging phù hợp cho trường hợp đơn giản, dữ liệu nhỏ, ít thay đổi.
- Cursor paging phù hợp cho hệ thống lớn, dữ liệu động, cần hiệu năng và độ ổn định cao.
