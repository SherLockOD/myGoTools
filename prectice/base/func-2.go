package main

import "time"

type Connection struct {
	addr string
}

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply (o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func Connect(addr string, opts ...Option) (*Connection, error) {
	defaultTimeout := time.Second
	defaultCaching := false

	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr: addr,
	}, nil
}

func main() {
	db := Co

}