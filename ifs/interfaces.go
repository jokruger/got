package ifs

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type Container[T any] interface {
	Contains(T) bool
}

type CompareProvider interface {
	Compare(CompareProvider) int
}

type ZeroTestProvider interface {
	IsZero() bool
}

type IDProvider[T any] interface {
	GetID() T
}

type NameProvider interface {
	GetName() string
}
