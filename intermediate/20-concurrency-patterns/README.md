# Panduan Lengkap Concurrency Patterns di Go

## 1. Pengenalan Concurrency di Go

### Apa itu Concurrency?

Concurrency adalah kemampuan untuk menjalankan beberapa tugas secara bersamaan (concurrent), bukan secara berurutan (sequential). Go dirancang khusus dengan concurrency sebagai fitur utama melalui konsep **goroutines** dan **channels**.

### Perbedaan Concurrency vs Parallelism

- **Concurrency**: Menangani banyak tugas secara bersamaan (dealing with lots of things at once)
- **Parallelism**: Melakukan banyak tugas secara bersamaan (doing lots of things at once)

Go focus pada concurrency, yang memungkinkan kode Anda untuk lebih efisien dalam menangani I/O operations, network requests, dan task-task yang memerlukan waiting time.

## 2. Goroutines - Building Block Concurrency

### Apa itu Goroutines?

Goroutines adalah lightweight threads yang dikelola oleh Go runtime. Berbeda dengan OS threads yang membutuhkan sekitar 2MB memory, goroutines hanya membutuhkan sekitar 2KB dan dapat dibuat jutaan tanpa masalah performance.

```go
package main

import (
    "fmt"
    "time"
)

// Fungsi biasa yang akan dijalankan sebagai goroutine
func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello from %s - %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Menjalankan fungsi secara sequential (berurutan)
    fmt.Println("=== Sequential Execution ===")
    sayHello("Alice")
    sayHello("Bob")
    
    fmt.Println("\n=== Concurrent Execution ===")
    // Menjalankan fungsi sebagai goroutine (concurrent)
    go sayHello("Charlie") // 'go' keyword membuat goroutine baru
    go sayHello("Dave")
    
    // Main function perlu menunggu goroutines selesai
    // Tanpa ini, program akan langsung exit
    time.Sleep(1 * time.Second)
    
    fmt.Println("Program finished")
}
```

**Penjelasan Kode:**
1. Function `sayHello` adalah function biasa yang mencetak pesan 3 kali
2. Pada sequential execution, Alice selesai dulu baru Bob dijalankan
3. Pada concurrent execution, Charlie dan Dave berjalan bersamaan
4. Keyword `go` sebelum function call membuat goroutine baru
5. `time.Sleep(1 * time.Second)` diperlukan agar main function menunggu goroutines selesai

## 3. Channels - Komunikasi antar Goroutines

### Konsep Dasar Channels

Channels adalah pipes yang menghubungkan concurrent goroutines. Anda bisa mengirim nilai dari satu goroutine dan menerima nilai tersebut dari goroutine lain melalui channel.

**Prinsip Go**: "Don't communicate by sharing memory; share memory by communicating"

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Membuat channel untuk mengirim string
    messages := make(chan string)
    
    // Goroutine yang mengirim pesan ke channel
    go func() {
        time.Sleep(500 * time.Millisecond)
        messages <- "Hello from goroutine!" // Mengirim ke channel
    }()
    
    // Menerima pesan dari channel (blocking operation)
    msg := <-messages // Menerima dari channel
    fmt.Println("Received:", msg)
    
    fmt.Println("Program finished")
}
```

**Penjelasan Kode:**
1. `make(chan string)` membuat channel yang bisa mengirim/menerima string
2. `messages <-` (send operation) mengirim nilai ke channel
3. `<- messages` (receive operation) menerima nilai dari channel
4. Channel operations bersifat **blocking** - goroutine akan menunggu sampai ada yang menerima/mengirim

### Buffered vs Unbuffered Channels

```go
package main

import (
    "fmt"
    "time"
)

func demonstrateUnbufferedChannel() {
    fmt.Println("=== Unbuffered Channel ===")
    ch := make(chan string) // Unbuffered channel
    
    go func() {
        fmt.Println("Goroutine: Sending message...")
        ch <- "Hello"
        fmt.Println("Goroutine: Message sent!")
    }()
    
    time.Sleep(100 * time.Millisecond) // Simulasi delay
    fmt.Println("Main: Receiving message...")
    msg := <-ch
    fmt.Println("Main: Received:", msg)
}

func demonstrateBufferedChannel() {
    fmt.Println("\n=== Buffered Channel ===")
    ch := make(chan string, 2) // Buffered channel dengan kapasitas 2
    
    go func() {
        fmt.Println("Goroutine: Sending first message...")
        ch <- "Hello"
        fmt.Println("Goroutine: First message sent!")
        
        fmt.Println("Goroutine: Sending second message...")
        ch <- "World"
        fmt.Println("Goroutine: Second message sent!")
    }()
    
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Main: Receiving first message...")
    msg1 := <-ch
    fmt.Println("Main: Received:", msg1)
    
    fmt.Println("Main: Receiving second message...")
    msg2 := <-ch
    fmt.Println("Main: Received:", msg2)
}

func main() {
    demonstrateUnbufferedChannel()
    demonstrateBufferedChannel()
}
```

**Perbedaan:**
- **Unbuffered Channel**: Send operation akan block sampai ada receiver
- **Buffered Channel**: Send operation tidak block sampai buffer penuh

## 4. Select Statement - Non-blocking Channel Operations

Select statement memungkinkan goroutine untuk menunggu multiple channel operations dan execute case yang pertama ready.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Membuat dua channel
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Goroutine yang mengirim ke ch1 setelah 1 detik
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from channel 1"
    }()
    
    // Goroutine yang mengirim ke ch2 setelah 500ms
    go func() {
        time.Sleep(500 * time.Millisecond)
        ch2 <- "Message from channel 2"
    }()
    
    // Select statement untuk menunggu channel mana yang ready duluan
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received from ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received from ch2:", msg2)
        case <-time.After(2 * time.Second): // Timeout case
            fmt.Println("Timeout!")
            return
        }
    }
}
```

**Penjelasan Select:**
1. Select akan menunggu sampai salah satu case ready untuk execute
2. Jika multiple case ready bersamaan, Go akan memilih secara random
3. `time.After()` berguna untuk membuat timeout
4. Select bersifat non-blocking jika ada `default` case

## 5. Common Concurrency Patterns

### Pattern 1: Worker Pool

Worker Pool adalah pattern untuk membatasi jumlah goroutines yang bekerja secara bersamaan. Sangat berguna untuk menghindari resource exhaustion.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Job represents work to be done
type Job struct {
    ID   int
    Data string
}

// Result represents the result of a job
type Result struct {
    JobID  int
    Output string
}

// Worker function that processes jobs
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done() // Menandakan worker selesai ketika function return
    
    for job := range jobs { // Iterasi semua jobs dari channel
        fmt.Printf("Worker %d starting job %d\n", id, job.ID)
        
        // Simulasi processing time
        time.Sleep(500 * time.Millisecond)
        
        // Kirim result ke channel results
        results <- Result{
            JobID:  job.ID,
            Output: fmt.Sprintf("Processed %s by worker %d", job.Data, id),
        }
        
        fmt.Printf("Worker %d finished job %d\n", id, job.ID)
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 10
    
    // Channels untuk jobs dan results
    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    
    // WaitGroup untuk menunggu semua workers selesai
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1) // Menambah counter untuk setiap worker
        go worker(i, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{
            ID:   j,
            Data: fmt.Sprintf("task-%d", j),
        }
    }
    close(jobs) // Tutup channel jobs (tidak ada job lagi)
    
    // Goroutine untuk menutup results channel setelah semua worker selesai
    go func() {
        wg.Wait()    // Tunggu semua worker selesai
        close(results) // Tutup channel results
    }()
    
    // Collect results
    fmt.Println("\n=== Results ===")
    for result := range results {
        fmt.Printf("Job %d: %s\n", result.JobID, result.Output)
    }
    
    fmt.Println("All jobs completed!")
}
```

**Keuntungan Worker Pool:**
1. **Resource Control**: Membatasi jumlah concurrent operations
2. **Better Performance**: Mencegah system overload
3. **Scalability**: Mudah adjust jumlah workers sesuai kebutuhan

### Pattern 2: Pipeline

Pipeline pattern memungkinkan data mengalir melalui serangkaian stages, dimana setiap stage diproses oleh goroutine yang berbeda.

```go
package main

import (
    "fmt"
    "sync"
)

// Stage 1: Generate numbers
func generateNumbers(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// Stage 2: Square numbers
func squareNumbers(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// Stage 3: Filter even numbers
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
    }()
    return out
}

func main() {
    // Input numbers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Printf("Input: %v\n", numbers)
    
    // Create pipeline
    numberChan := generateNumbers(numbers...)
    squaredChan := squareNumbers(numberChan)
    evenChan := filterEven(squaredChan)
    
    // Consume results
    fmt.Print("Pipeline result (squared even numbers): ")
    var results []int
    for result := range evenChan {
        results = append(results, result)
    }
    fmt.Printf("%v\n", results)
}
```

**Pipeline dengan Fan-Out Fan-In Pattern:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Expensive computation simulation
func expensiveComputation(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            // Simulasi computation yang berat
            time.Sleep(100 * time.Millisecond)
            out <- n * n * n // Cubic
        }
    }()
    return out
}

// Fan-Out: Distribute work to multiple goroutines
func fanOut(in <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        outputs[i] = expensiveComputation(in)
    }
    return outputs
}

// Fan-In: Merge multiple channels into one
func fanIn(inputs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    // Start a goroutine for each input channel
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                out <- n
            }
        }(input)
    }
    
    // Close output channel when all inputs are done
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

func main() {
    // Generate input
    input := make(chan int)
    go func() {
        defer close(input)
        for i := 1; i <= 10; i++ {
            input <- i
        }
    }()
    
    // Fan-out to 3 workers
    fmt.Println("Processing with Fan-Out Fan-In pattern...")
    start := time.Now()
    
    workers := fanOut(input, 3)
    result := fanIn(workers...)
    
    // Collect results
    var results []int
    for r := range result {
        results = append(results, r)
    }
    
    fmt.Printf("Results: %v\n", results)
    fmt.Printf("Processing time: %v\n", time.Since(start))
}
```

### Pattern 3: Rate Limiting

Rate limiting berguna untuk membatasi frequency of operations, misalnya API calls.

```go
package main

import (
    "fmt"
    "time"
)

// Simple rate limiter using channel
func simpleRateLimiter() {
    fmt.Println("=== Simple Rate Limiter ===")
    
    // Create a rate limiter that allows 1 request per 200ms
    rateLimiter := time.Tick(200 * time.Millisecond)
    
    // Simulate 5 requests
    for i := 1; i <= 5; i++ {
        <-rateLimiter // Wait for rate limiter
        fmt.Printf("Request %d processed at %s\n", i, time.Now().Format("15:04:05.000"))
    }
}

// Burst rate limiter using buffered channel
func burstRateLimiter() {
    fmt.Println("\n=== Burst Rate Limiter ===")
    
    // Create a burst limiter that allows 3 requests immediately, then 1 per second
    burstLimiter := make(chan time.Time, 3)
    
    // Fill up the burst limiter initially
    for i := 0; i < 3; i++ {
        burstLimiter <- time.Now()
    }
    
    // Refill the burst limiter every second
    go func() {
        for t := range time.Tick(1 * time.Second) {
            select {
            case burstLimiter <- t:
                // Token added
            default:
                // Channel is full, drop the token
            }
        }
    }()
    
    // Simulate 7 requests
    for i := 1; i <= 7; i++ {
        <-burstLimiter
        fmt.Printf("Burst request %d processed at %s\n", i, time.Now().Format("15:04:05.000"))
    }
}

// Token bucket rate limiter
type TokenBucket struct {
    tokens chan struct{}
    ticker *time.Ticker
}

func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
    tb := &TokenBucket{
        tokens: make(chan struct{}, capacity),
        ticker: time.NewTicker(refillRate),
    }
    
    // Fill initial tokens
    for i := 0; i < capacity; i++ {
        select {
        case tb.tokens <- struct{}{}:
        default:
        }
    }
    
    // Start refill goroutine
    go tb.refill()
    
    return tb
}

func (tb *TokenBucket) refill() {
    for range tb.ticker.C {
        select {
        case tb.tokens <- struct{}{}:
            // Token added
        default:
            // Bucket is full
        }
    }
}

func (tb *TokenBucket) Allow() bool {
    select {
    case <-tb.tokens:
        return true
    default:
        return false
    }
}

func (tb *TokenBucket) Wait() {
    <-tb.tokens
}

func (tb *TokenBucket) Close() {
    tb.ticker.Stop()
}

func tokenBucketExample() {
    fmt.Println("\n=== Token Bucket Rate Limiter ===")
    
    // Create token bucket: 2 tokens capacity, refill 1 token every 500ms
    bucket := NewTokenBucket(2, 500*time.Millisecond)
    defer bucket.Close()
    
    // Simulate requests
    for i := 1; i <= 6; i++ {
        if bucket.Allow() {
            fmt.Printf("Token bucket request %d: ALLOWED at %s\n", i, time.Now().Format("15:04:05.000"))
        } else {
            fmt.Printf("Token bucket request %d: DENIED at %s\n", i, time.Now().Format("15:04:05.000"))
        }
        time.Sleep(200 * time.Millisecond) // Request interval
    }
    
    fmt.Println("\nWaiting for tokens to refill...")
    time.Sleep(1 * time.Second)
    
    // Try again after waiting
    if bucket.Allow() {
        fmt.Printf("After waiting: ALLOWED at %s\n", time.Now().Format("15:04:05.000"))
    }
}

func main() {
    simpleRateLimiter()
    burstRateLimiter()
    tokenBucketExample()
}
```

## 6. Error Handling dalam Concurrent Programs

### Pattern untuk Error Handling

```go
package main

import (
    "errors"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Result dengan error handling
type TaskResult struct {
    ID     int
    Result string
    Error  error
}

// Task yang bisa gagal
func unreliableTask(id int) TaskResult {
    // Simulasi random failure
    if rand.Float32() < 0.3 { // 30% chance of failure
        return TaskResult{
            ID:    id,
            Error: fmt.Errorf("task %d failed randomly", id),
        }
    }
    
    // Simulasi processing time
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    
    return TaskResult{
        ID:     id,
        Result: fmt.Sprintf("Task %d completed successfully", id),
    }
}

// Worker dengan error handling
func workerWithErrorHandling(id int, jobs <-chan int, results chan<- TaskResult, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for jobID := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, jobID)
        result := unreliableTask(jobID)
        results <- result
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    results := make(chan TaskResult, numJobs)
    
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go workerWithErrorHandling(i, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Close results channel after all workers done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results and handle errors
    var successful []TaskResult
    var failed []TaskResult
    
    for result := range results {
        if result.Error != nil {
            failed = append(failed, result)
            fmt.Printf("❌ Job %d failed: %v\n", result.ID, result.Error)
        } else {
            successful = append(successful, result)
            fmt.Printf("✅ Job %d: %s\n", result.ID, result.Result)
        }
    }
    
    fmt.Printf("\n=== Summary ===\n")
    fmt.Printf("Successful: %d\n", len(successful))
    fmt.Printf("Failed: %d\n", len(failed))
}
```

## 7. Context untuk Cancellation dan Timeout

Context adalah cara standar di Go untuk menangani cancellation, timeout, dan passing values across API boundaries.

```go
package main

import (
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Long running task yang respect context
func longRunningTask(ctx context.Context, id int, results chan<- string) {
    // Simulasi task yang membutuhkan waktu lama
    workDuration := time.Duration(rand.Intn(3000)+1000) * time.Millisecond
    
    fmt.Printf("Task %d starting (estimated duration: %v)\n", id, workDuration)
    
    select {
    case <-time.After(workDuration):
        // Task completed normally
        results <- fmt.Sprintf("Task %d completed successfully", id)
    case <-ctx.Done():
        // Task cancelled
        results <- fmt.Sprintf("Task %d cancelled: %v", id, ctx.Err())
    }
}

func demonstrateTimeout() {
    fmt.Println("=== Timeout Example ===")
    
    // Create context dengan timeout 2 detik
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Penting: selalu call cancel untuk cleanup
    
    results := make(chan string, 5)
    var wg sync.WaitGroup
    
    // Start 5 tasks
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(taskID int) {
            defer wg.Done()
            longRunningTask(ctx, taskID, results)
        }(i)
    }
    
    // Wait for all tasks and close results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}

func demonstrateCancellation() {
    fmt.Println("\n=== Manual Cancellation Example ===")
    
    ctx, cancel := context.WithCancel(context.Background())
    
    results := make(chan string, 3)
    var wg sync.WaitGroup
    
    // Start 3 tasks
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(taskID int) {
            defer wg.Done()
            longRunningTask(ctx, taskID, results)
        }(i)
    }
    
    // Cancel after 1.5 seconds
    go func() {
        time.Sleep(1500 * time.Millisecond)
        fmt.Println("Cancelling all tasks...")
        cancel()
    }()
    
    go func() {
        wg.Wait()
        close(results)
    }()
    
    for result := range results {
        fmt.Println(result)
    }
}

// HTTP request simulation with context
func simulateHTTPRequest(ctx context.Context, url string, results chan<- string) {
    // Simulasi network request
    requestTime := time.Duration(rand.Intn(2000)+500) * time.Millisecond
    
    select {
    case <-time.After(requestTime):
        results <- fmt.Sprintf("✅ Successfully fetched %s (took %v)", url, requestTime)
    case <-ctx.Done():
        results <- fmt.Sprintf("❌ Request to %s cancelled: %v", url, ctx.Err())
    }
}

func demonstrateHTTPWithTimeout() {
    fmt.Println("\n=== HTTP Requests with Timeout ===")
    
    urls := []string{
        "https://api1.example.com",
        "https://api2.example.com",
        "https://api3.example.com",
        "https://api4.example.com",
    }
    
    // Context dengan timeout 1 detik
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    results := make(chan string, len(urls))
    var wg sync.WaitGroup
    
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            simulateHTTPRequest(ctx, u, results)
        }(url)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
    
    for result := range results {
        fmt.Println(result)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    demonstrateTimeout()
    demonstrateCancellation()
    demonstrateHTTPWithTimeout()
}
```

## 8. Best Practices dan Common Pitfalls

### Best Practices

1. **Selalu close channels yang Anda buat**
2. **Gunakan context untuk cancellation dan timeout**
3. **Hindari sharing memory, gunakan channels**
4. **Gunakan sync.WaitGroup untuk menunggu multiple goroutines**
5. **Handle errors dengan proper**
6. **Gunakan buffered channels untuk performance ketika diperlukan**

### Common Pitfalls

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// ❌ PITFALL 1: Goroutine leak karena tidak ada receiver
func goroutineLeakExample() {
    fmt.Println("=== Goroutine Leak Example ===")
    
    ch := make(chan string)
    
    // Goroutine ini akan stuck selamanya karena tidak ada receiver
    go func() {
        ch <- "This will block forever"
        fmt.Println("This will never print")
    }()
    
    // Main function tidak membaca dari channel
    fmt.Println("Main function exits, goroutine leaks")
    time.Sleep(100 * time.Millisecond)
}

// ✅ SOLUTION: Selalu pastikan ada receiver atau gunakan buffered channel
func goroutineLeakSolution() {
    fmt.Println("\n=== Goroutine Leak Solution ===")
    
    ch := make(chan string, 1) // Buffered channel
    
    go func() {
        ch <- "This won't block"
        fmt.Println("Goroutine completed successfully")
    }()
    
    // Atau pastikan ada receiver:
    // msg := <-ch
    // fmt.Println("Received:", msg)
    
    time.Sleep(100 * time.Millisecond)
}

// ❌ PITFALL 2: Race condition
func raceConditionExample() {
    fmt.Println("\n=== Race Condition Example ===")
    
    counter := 0
    var wg sync.WaitGroup
    
    // Multiple goroutines mengakses variable yang sama
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                counter++ // Race condition!
            }
        }()
    }
    
    wg.Wait()
    fmt.Printf("Expected: 10000, Got: %d\n", counter)
}

// ✅ SOLUTION: Gunakan sync.Mutex atau channels
func raceConditionSolution() {
    fmt.Println("\n=== Race Condition Solution ===")
    
    counter := 0
    var mu sync.Mutex
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                mu.Lock()
                counter++
                mu.Unlock()
            }
        }()
    }
    
    wg.Wait()
    fmt.Printf("Expected: 10000, Got: %d\n", counter)
}

// ❌ PITFALL 3: Deadlock
func deadlockExample() {
    fmt.Println("\n=== Deadlock Example ===")
    
    ch := make(chan string)
    
    // This will cause deadlock because no goroutine is reading
    // Uncomment to see deadlock:
    // ch <- "This will cause deadlock"
    
    fmt.Println("Deadlock avoided by not executing the problematic code")
}

// ✅ SOLUTION: Pastikan ada balance antara sender dan receiver
func deadlockSolution() {
    fmt.Println("\n=== Deadlock Solution ===")
    
    ch := make(chan string, 1) // Buffered channel
    
    ch <- "This won't cause deadlock"
    msg := <-ch
    fmt.Println("Received:", msg)