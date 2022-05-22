package yeh

import "fmt"

func MustWith2[T1, T2 any](t1 T1, t2 T2, err error) mustWith2[T1, T2] {
	return mustWith2[T1, T2]{
		t1:  t1,
		t2:  t2,
		err: err,
	}
}

type mustWith2[T1, T2 any] struct {
	t1  T1
	t2  T2
	err error
}

func (m mustWith2[T1, T2]) Replace(err error) (T1, T2) {
	if m.err != nil {
		panic(err)
	}
	return m.t1, m.t2
}

func (m mustWith2[T1, T2]) Wrap(err error) (T1, T2) {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
	return m.t1, m.t2
}

func (m mustWith2[T1, T2]) Callback(fn FnErrorCallback) (T1, T2) {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
	return m.t1, m.t2
}
