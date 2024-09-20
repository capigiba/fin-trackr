# DANH SÁCH CÁC TÍNH NĂNG CẦN LÀM

* [Dịch sang  tiếng anh](../README.md)


## 1. Ứng Dụng Theo Dõi Chi Tiêu
Mô tả: Xây dựng một ứng dụng đơn giản để theo dõi chi tiêu cá nhân. Người dùng có thể thêm, chỉnh sửa, và xóa các giao dịch, phân loại chúng và xem các báo cáo cơ bản (ví dụ: tổng chi tiêu theo tháng).
Các khái niệm chính:
Làm việc với các thao tác CRUD và cơ sở dữ liệu.
Hiểu cách giao dịch tài chính và phân loại chúng.
Hình dung dữ liệu cơ bản (tùy chọn).
Các khía cạnh học tập:

Xử lý cơ sở dữ liệu trong Golang (ví dụ: sử dụng SQLite hoặc PostgreSQL).
Triển khai REST APIs cho các giao dịch tài chính.
Bảo mật khi xử lý dữ liệu tài chính cá nhân.

## 2. API Chuyển Đổi Tiền Tệ
Mô tả: Tạo một API chuyển đổi tiền tệ lấy tỷ giá hối đoái theo thời gian thực từ một API bên ngoài và thực hiện chuyển đổi tiền tệ cho người dùng.
Các khái niệm chính:
Làm việc với các API tài chính bên ngoài (ví dụ: OpenExchangeRates).
Xử lý và tính toán dữ liệu.
Các khía cạnh học tập:

Sử dụng HTTP client của Golang để gọi các API bên ngoài.
Triển khai giới hạn tốc độ cho các API.
Chuyển đổi tiền tệ và quản lý dữ liệu tỷ giá hối đoái.

## 3. Cổng Thanh Toán Đơn Giản
Mô tả: Xây dựng một cổng thanh toán đơn giản cho phép người dùng thực hiện thanh toán (giả lập chi tiết thẻ tín dụng) và xử lý các giao dịch, tạo biên lai thanh toán.
Các khái niệm chính:
Xác thực và xử lý thanh toán.
Giao dịch giả lập để mô phỏng cách hoạt động của một cổng thanh toán.
Các khía cạnh học tập:

Sử dụng JSON và HTTP cho các yêu cầu.
Hiểu cách các cổng thanh toán thực sự hoạt động.
Bảo mật khi xử lý dữ liệu thanh toán.

## 4. Ứng Dụng Lập Ngân Sách
Mô tả: Tạo một công cụ lập ngân sách cho phép người dùng đặt ngân sách cho các danh mục khác nhau (ví dụ: mua sắm, giải trí) và theo dõi chi tiêu của họ so với các ngân sách đó.
Các khái niệm chính:
Tạo và theo dõi ngân sách.
Xử lý các danh mục tài chính khác nhau.
Các khía cạnh học tập:

Triển khai logic nghiệp vụ xung quanh ngân sách.
Làm việc với dữ liệu tài chính theo thời gian (ví dụ: ngân sách hàng tháng).
Sử dụng các cấu trúc trong Golang và các tính năng xác thực dữ liệu.

## 5. Mô Phỏng Khoản Vay Nhỏ
Mô tả: Xây dựng một công cụ mô phỏng khoản vay cơ bản tính toán lịch trả nợ cho người dùng dựa trên số tiền vay, lãi suất và thời hạn khoản vay.
Các khái niệm chính:
Hiểu về lịch trả nợ vay.
Tính toán lãi suất (lãi kép, lịch trả nợ theo kiểu trả dần).
Các khía cạnh học tập:

Xử lý các tính toán tài chính trong Golang.
Tạo và quản lý dữ liệu khoản vay (ví dụ: sử dụng cơ sở dữ liệu SQL).
Mô phỏng và hình dung các khoản thanh toán theo thời gian.

## 6. Công Cụ Theo Dõi Danh Mục Cổ Phiếu
Mô tả: Ứng dụng cho phép người dùng theo dõi giá trị danh mục đầu tư cổ phiếu của họ, bao gồm cập nhật giá cổ phiếu theo thời gian thực từ một API công cộng.
Các khái niệm chính:
Làm việc với dữ liệu tài chính theo thời gian thực.
Quản lý danh mục đầu tư và tính toán giá trị danh mục.
Các khía cạnh học tập:

Lấy giá cổ phiếu từ một API bên ngoài (ví dụ: Alpha Vantage).
Làm việc với dữ liệu chuỗi thời gian trong Golang.
Xử lý và cập nhật dữ liệu theo thời gian thực.

## 7. Hệ Thống Xác Minh KYC (Know Your Customer)
Mô tả: Xây dựng hệ thống KYC đơn giản cho phép người dùng nộp tài liệu (ví dụ: CMND, hộ chiếu) và hệ thống xác minh và lưu trữ thông tin của họ.
Các khái niệm chính:
Xác minh và quản lý tài liệu.
Xử lý dữ liệu khách hàng nhạy cảm một cách bảo mật.
Các khía cạnh học tập:

Tải lên tệp và lưu trữ an toàn trong Golang.
Xác thực dữ liệu và cơ chế xác thực.
Hiểu các yêu cầu KYC đối với các hệ thống tài chính.

## 8. Công Cụ Theo Dõi Giá Tiền Điện Tử
Mô tả: Tạo ứng dụng lấy và hiển thị giá tiền điện tử (ví dụ: Bitcoin, Ethereum) từ một API công cộng và hiển thị xu hướng theo thời gian.
Các khái niệm chính:
Lấy dữ liệu tài chính theo thời gian thực.
Hình dung dữ liệu giá theo xu hướng.
Các khía cạnh học tập:

Sử dụng gói HTTP của Golang để gọi API.
Hiểu cách hoạt động của các API tiền điện tử.
Xử lý và hình dung dữ liệu chuỗi thời gian.

## 9. Hệ Thống Tạo Hóa Đơn Đơn Giản
Mô tả: Phát triển một công cụ tạo hóa đơn cho các dịch vụ/sản phẩm, bao gồm thông tin khách hàng, chi phí chi tiết và tổng số tiền phải trả.
Các khái niệm chính:
Tạo hóa đơn và tạo tệp PDF.
Quản lý dữ liệu khách hàng và sản phẩm/dịch vụ.
Các khía cạnh học tập:

Tạo PDF trong Golang (ví dụ: sử dụng gofpdf).
Xử lý logic nghiệp vụ liên quan đến hóa đơn và điều khoản thanh toán.
Thao tác cơ sở dữ liệu để lưu trữ dữ liệu hóa đơn và khách hàng.

## 10. Công Cụ Đối Chiếu Giao Dịch
Mô tả: Xây dựng một công cụ so sánh sao kê ngân hàng với hồ sơ giao dịch nội bộ để xác định sự sai lệch.
Các khái niệm chính:
Phân tích và so sánh sao kê ngân hàng.
Thuật toán đối chiếu để khớp các giao dịch.
Các khía cạnh học tập:

Phân tích tệp (CSV, XLSX, v.v.) trong Golang.
Thuật toán để đối chiếu tập dữ liệu lớn.
Tự động hóa kiểm tra và cân đối tài chính.
