package yeh

import "fmt"

func MustWith4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) mustWith4[T1, T2, T3, T4] {
	return mustWith4[T1, T2, T3, T4]{
		t1:  t1,
		t2:  t2,
		t3:  t3,
		t4:  t4,
		err: err,
	}
}

type mustWith4[T1, T2, T3, T4 any] struct {
	t1  T1
	t2  T2
	t3  T3
	t4  T4
	err error
}

func (m mustWith4[T1, T2, T3, T4]) Replace(err error) (T1, T2, T3, T4) {
	if m.err != nil {
		panic(err)
	}
	return m.t1, m.t2, m.t3, m.t4
}

func (m mustWith4[T1, T2, T3, T4]) Wrap(err error) (T1, T2, T3, T4) {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
	return m.t1, m.t2, m.t3, m.t4
}

func (m mustWith4[T1, T2, T3, T4]) Callback(fn FnErrorCallback) (T1, T2, T3, T4) {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
	return m.t1, m.t2, m.t3, m.t4
}
