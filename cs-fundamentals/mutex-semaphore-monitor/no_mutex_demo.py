import threading
import time

# Shared Resource
shared_counter = 0

# Note: No Mutex is used here

def increment_counter(thread_name):
    global shared_counter

    print(f"[{thread_name}] Starting task (No Lock)...")

    # CRITICAL SECTION (Unprotected)
    # Simulate a time-consuming task (Reading -> Processing -> Writing)
    local_copy = shared_counter
    
    # Simulate processing time (Context switch happens here)
    time.sleep(2) 
    
    # Write back the new value
    shared_counter = local_copy + 1

    print(f"[{thread_name}] Finished. Local value was {local_copy}, wrote {shared_counter}")

# Initialize threads
threads = []
print("--- Starting Simulation WITHOUT Mutex (Race Condition) ---")
for i in range(5):
    t = threading.Thread(target=increment_counter, args=(f"Thread-{i+1}",))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

print("-" * 30)
print(f"Expected result: 5")
print(f"Actual result:   {shared_counter}") 
print("-" * 30)