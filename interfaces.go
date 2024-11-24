package got

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// Container is a generic interface for a container that can check if it contains an item.
type Container[T any] interface {
	Contains(item T) bool
}

// Adder is a generic interface for a container that can add an item.
type Adder[T any] interface {
	Add(item T)
}

// Remover is a generic interface for a container that can remove an item.
type Remover[T any] interface {
	Remove(item T)
}

// MutableContainer is a generic interface for a container that can add, remove, and check if it contains an item.
type MutableContainer[T any] interface {
	Adder[T]
	Remover[T]
	Container[T]
}
