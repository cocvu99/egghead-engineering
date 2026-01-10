# Explain synchronization primitives: Mutexes, Semaphores, Monitors.

A. Mutex (Mutual Exclusion)

- Định nghĩa: Là một cái khóa
- Cơ chế hoạt động: "Chỉ 1 người được vào". Nếu Thread A đang giữ Mutex, Thead B muốn vào phải đứng chờ cho đến khi A nhả khóa ra
- Ví dụ: Trong toilet trên máy bay. Chỉ một người vào được, người đó chốt cửa (Lock). Người bên ngoài thấy đèn đỏ thì phải đợi (Wait)
- Ownership (Tính sở hữu - Rất quan trọng): Mutex có khái niệm ownership. Thread nào khóa (lock) thì chính Thread đó phải mở (lock). Thread A khóa mà Thread B mở -> Lỗi (Exception)
- Trạng thái Mutex: Chỉ có 2 trạng thái là Locked (đang bị giữ) và Unlocked 