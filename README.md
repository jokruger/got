# Overview
Package GOT implements frequently used functions and algorights.

## Installation

```bash
go get github.com/jokruger/got
```

## Batches

### func GoInBatches
```go
func GoInBatches[T any](tasks []T, batchSize int, f func([]T) error) error
```
GoInBatches runs a function in batches of tasks.

## Ranges

### func Range
```go
func Range(n int) []int
```
Range returns a slice of integers from 0 to n-1.
