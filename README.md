# Overview
Package GOT implements frequently used functions and algorights.

## Installation

```bash
go get github.com/jokruger/got
```


## Batches

### func Batches
```go
func Batches[T any](tasks []T, batchSize int) iter.Seq[[]T]
```
Batches returns an iterator that yields batches of tasks.

### func GoInBatches
```go
func GoInBatches[T any](tasks []T, batchSize int, f func([]T) error) error
```
GoInBatches runs a function in batches of tasks.


## Filters

## func Filter
```go
func Filter[T any](is []T, f func(T) bool) []T 
```
Filter returns a new slice containing only the elements of slice 'is' that satisfy the predicate f.

## func FilterIter
```go
func FilterIter[T any](is iter.Seq[T], f func(T) bool) []T
```
FilterIter returns a new slice containing only the elements of sequence 'is' that satisfy the predicate f.


## Maps

### func Map
```go
func Map[T any, U any](is []T, f func(T) U) []U
```
Map applies a function to each element of a slice and returns a new slice.

### func MapIter
```go
func MapIter[T any, U any](is iter.Seq[T], f func(T) U) []U
```
MapIter applies a function to each element of a sequence and returns a new slice.

### func MapUnique
```go
func MapUnique[T any, U comparable](is []T, f func(T) U) []U
```
MapUnique applies a function to each element of a slice and returns a new slice with unique elements.

### func MapUniqueIter
```go
func MapUniqueIter[T any, U comparable](is iter.Seq[T], f func(T) U) []U
```
MapUniqueIter applies a function to each element of a sequence and returns a new slice with unique elements.


## Partitioning

### func Partitions
```go
func Partitions[T any](items []T, eq func(a, b T) bool) iter.Seq[[]T]
```
Partitions returns a sequence of partitions of the input items based on 'eq' comparator.


## Ranges

### func Range
```go
func Range(n int) []int
```
Range returns a slice of integers from 0 to n-1.


## Other

### func Choose
```go
func Choose[T comparable](values ...T) T
```
Choose returns the first non-zero value from the list of values.
