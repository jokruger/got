# Overview
Package GOT implements frequently used functions and algorights.

## Installation

```bash
go get github.com/jokruger/got
```


## Chunks

### func Chunks
```go
func Chunks[T any](ts []T, size int) iter.Seq[[]T]
```
Chunks returns a sequence of chunks of the input slice, each of given size at most.

### func GoInChunks
```go
func GoInChunks[T any](ts []T, size int, f func([]T) error) error
```
GoInChunks calls the given function for each chunk of the input slice, each of given size at most.


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

## func FilterSet
```go
func FilterSet[T comparable](is []T, s Set[T]) []T 
```
FilterSet returns a new slice containing only the elements of slice 'is' that are in set 's'.

## func FilterSetIter
```go
func FilterSetIter[T comparable](is iter.Seq[T], s Set[T]) []T
```
FilterSetIter returns a new slice containing only the elements of sequence 'is' that are in set 's'.


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


## Sets

### type Set
Set is a set of elements of type T.

```go
func NewSet[T comparable](src ...Set[T]) Set[T]
```
NewSet creates a new set. The set can be initialized with other sets.

```go
func NewSetFromSlice[T comparable](s []T) Set[T]
```
NewSetFromSlice creates a new set from a slice.

```go
func NewSetFromIter[T comparable](is iter.Seq[T]) Set[T]
```
NewSetFromIter creates a new set from an iterator.

```go
func (s Set[T]) Len() int
```
Len returns the number of elements in the set.

```go
func (s Set[T]) Add(es ...T)
```
Add adds elements to the set.

```go
func (s Set[T]) AddIter(is iter.Seq[T])
```
AddIter adds elements to the set from an iterator.

```go
func (s Set[T]) AddSet(other Set[T])
```
AddSet adds elements to the set from another set.

```go
func (s Set[T]) AddSlice(es []T)
```
AddSlice adds elements to the set from a slice.

```go
func (s Set[T]) Remove(es ...T)
```
Remove removes elements from the set.

```go
func (s Set[T]) Contains(e T) bool
```
Contains returns true if the set contains the element.

```go
func (s Set[T]) ContainsAll(es ...T) bool
```
ContainsAll returns true if the set contains all the elements.

```go
func (s Set[T]) ContainsAny(es ...T) bool
```
ContainsAny returns true if the set contains any of the elements.

```go
func (s Set[T]) Slice() []T 
```
Slice returns the elements of the set as a slice.

```go
func (s Set[T]) Iter() iter.Seq[T]
```
Iter returns an iterator over the elements of the set.


## Other

### func Choose
```go
func Choose[T comparable](values ...T) T
```
Choose returns the first non-zero value from the list of values.
