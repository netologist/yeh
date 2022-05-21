package yeh

import (
	"fmt"
)

type FnErrorCallback func(error) error

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Must2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

func Must3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3
}

func Must4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4
}

func Must5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5
}

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
