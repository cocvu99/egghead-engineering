# Một thử thách nhỏ (Socratic method): Sau khi chạy thử file no_mutex_demo.py, hãy so sánh Tổng thời gian chạy của 2 file.

- File mutex_demo.py chạy rất lâu (Tổng thời gian = 5 thread x thời gian sleep).

- File no_mutex_demo.py chạy rất nhanh (Tổng thời gian = 1 lần thời gian sleep). -> Trong thực tế, làm sao để chúng ta vừa đảm bảo dữ liệu đúng (như file 1) mà lại muốn chạy nhanh (như file 2)? 

(Gợi ý: Database thực tế hiếm khi lock toàn bộ bảng như thế này).