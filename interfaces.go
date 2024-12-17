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

// Identifiable defines types that implements the GetID method.
type Identifiable[T any] interface {
	GetID() T
}

// ProductIDProvider defines types that implements the GetProductID method.
type ProductIDProvider[T any] interface {
	GetProductID() T
}

// CustomerIDProvider defines types that implements the GetCustomerID method.
type CustomerIDProvider[T any] interface {
	GetCustomerID() T
}

// EventIDProvider defines types that implements the GetEventID method.
type EventIDProvider[T any] interface {
	GetEventID() T
}

// AccountIDProvider defines types that implements the GetAccountID method.
type AccountIDProvider[T any] interface {
	GetAccountID() T
}

// Named defines types that implements the GetName method.
type Named[T any] interface {
	GetName() T
}

// EventNameProvider defines types that implements the GetEventName method.
type EventNameProvider[T any] interface {
	GetEventName() T
}

// ValueTimeProvider defines types that implements the GetValueTime method.
type ValueTimeProvider[T any] interface {
	GetValueTime() T
}

// CreatedAtProvider defines types that implements the GetCreatedAt method.
type CreatedAtProvider[T any] interface {
	GetCreatedAt() T
}

// UpdatedAtProvider defines types that implements the GetUpdatedAt method.
type UpdatedAtProvider[T any] interface {
	GetUpdatedAt() T
}

// AmountProvider defines types that implements the GetAmount method.
type AmountProvider[T any] interface {
	GetAmount() T
}

// TransactionIDProvider defines types that implements the GetTransactionID method.
type TransactionIDProvider[T any] interface {
	GetTransactionID() T
}

// PriorityProvider defines types that implements the GetPriority method.
type PriorityProvider[T any] interface {
	GetPriority() T
}
