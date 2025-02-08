```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var counter int
        var mu sync.Mutex // Add a mutex for synchronization
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Acquire the lock before accessing the counter
                        counter++
                        mu.Unlock() // Release the lock after incrementing
                }()
        }

        wg.Wait()
        fmt.Println("Counter value:", counter)
}
```