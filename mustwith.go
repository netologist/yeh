package yeh

import "fmt"

type FnErrorCallback func(error) error

func MustWith[T any](t T, err error) mustWith[T] {
	return mustWith[T]{
		t:   t,
		err: err,
	}
}

type mustWith[T any] struct {
	t   T
	err error
}

func (m mustWith[T]) Replace(err error) T {
	if m.err != nil {
		panic(err)
	}
	return m.t
}

func (m mustWith[T]) Wrap(err error) T {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
	return m.t
}

func (m mustWith[T]) Callback(fn FnErrorCallback) T {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
	return m.t
}
