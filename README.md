# Go Race Condition Example

This repository demonstrates a simple race condition in Go that can lead to unexpected results when multiple goroutines access and modify a shared variable concurrently.

## The Problem
The `bug.go` file contains a program that launches multiple goroutines to increment a counter. Without proper synchronization mechanisms, the increment operations can overlap, resulting in a final counter value lower than expected.

## The Solution
The `bugSolution.go` file provides a corrected version of the program. It uses a mutex to protect the counter from concurrent access, ensuring that each increment operation is atomic and produces the correct final count.

## How to Run
1. Clone the repository.
2. Navigate to the repository's directory.
3. Run the buggy code: `go run bug.go`
4. Run the corrected code: `go run bugSolution.go`

Observe the difference in the counter's final value.