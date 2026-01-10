import threading
import time

# Shared Resource
shared_counter = 0

# Initialize the Mutex (Lock)
mutex = threading.Lock()

def increment_counter(thread_name):
    global shared_counter

    print(f"[{thread_name}] Waiting for Mutex...")

    # Critical Section: Only one thread can enter this block at a time
    with mutex:
        print(f"[{thread_name}] Acquired Mutex. Processing...")

        # Simulate a time-consuming task to visualize blocking mechanism
        local_copy = shared_counter
        time.sleep(2) # 2 seconds
        shared_counter = local_copy + 1

        print(f"[{thread_name}] Updated counter to {shared_counter}")

    # End of Critical Section (Mutex is automatically released here)

# Initialize threads
threads = []
print("--- Starting Simulation WITH Mutex ---")
for i in range(5):
    t = threading.Thread(target=increment_counter, args=(f"Thread-{i+1}",))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

print(f"Final result: {shared_counter}")