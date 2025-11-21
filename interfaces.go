package got

// Integer defines numeric types that are integers (signed or unsigned).
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Numeric defines numeric types that can be either integers or floating-point numbers.
type Numeric interface {
	Integer | ~float32 | ~float64
}

// Ordered defines types that can be ordered.
type Ordered interface {
	Integer | ~float32 | ~float64 | ~string
}

// Container defines types that implements the Contains method.
type Container[T any] interface {
	Contains(T) bool
}

// Comparable defines types that implements the Compare method.
type Comparable[T any] interface {
	Compare(T) int
}

// ZeroCheckable defines types that implements the IsZero method.
type ZeroCheckable interface {
	IsZero() bool
}

// Identifiable defines types that implements the ID method.
type Identifiable[T any] interface {
	ID() T
}

// ProductIDProvider defines types that implements the ProductID method.
type ProductIDProvider[T any] interface {
	ProductID() T
}

// CustomerIDProvider defines types that implements the CustomerID method.
type CustomerIDProvider[T any] interface {
	CustomerID() T
}

// EventIDProvider defines types that implements the EventID method.
type EventIDProvider[T any] interface {
	EventID() T
}

// AccountIDProvider defines types that implements the AccountID method.
type AccountIDProvider[T any] interface {
	AccountID() T
}

// Named defines types that implements the Name method.
type Named[T any] interface {
	Name() T
}

// EventNameProvider defines types that implements the EventName method.
type EventNameProvider[T any] interface {
	EventName() T
}

// ValueTimeProvider defines types that implements the ValueTime method.
type ValueTimeProvider[T any] interface {
	ValueTime() T
}

// CreatedAtProvider defines types that implements the CreatedAt method.
type CreatedAtProvider[T any] interface {
	CreatedAt() T
}

// UpdatedAtProvider defines types that implements the UpdatedAt method.
type UpdatedAtProvider[T any] interface {
	UpdatedAt() T
}

// AmountProvider defines types that implements the Amount method.
type AmountProvider[T any] interface {
	Amount() T
}

// TransactionIDProvider defines types that implements the TransactionID method.
type TransactionIDProvider[T any] interface {
	TransactionID() T
}

// PriorityProvider defines types that implements the Priority method.
type PriorityProvider[T any] interface {
	Priority() T
}
