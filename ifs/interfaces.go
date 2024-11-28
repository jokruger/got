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

type CompareProvider[T any] interface {
	Compare(T) int
}

type ZeroTestProvider interface {
	IsZero() bool
}

type IDProvider[T any] interface {
	GetID() T
}

type ProductIDProvider[T any] interface {
	GetProductID() T
}

type CustomerIDProvider[T any] interface {
	GetCustomerID() T
}

type EventIDProvider[T any] interface {
	GetEventID() T
}

type AccountIDProvider[T any] interface {
	GetAccountID() T
}

type NameProvider[T any] interface {
	GetName() T
}

type ValueTimeProvider[T any] interface {
	GetValueTime() T
}

type CreatedAtProvider[T any] interface {
	GetCreatedAt() T
}

type UpdatedAtProvider[T any] interface {
	GetUpdatedAt() T
}
