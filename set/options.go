package set

type Options struct {
	capacity int
}

type Option func(*Options)

func WithCapacity(capacity int) Option {
	return func(o *Options) {
		o.capacity = capacity
	}
}
