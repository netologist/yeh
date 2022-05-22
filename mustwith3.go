package yeh

import "fmt"

func MustWith3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) mustWith3[T1, T2, T3] {
	return mustWith3[T1, T2, T3]{
		t1:  t1,
		t2:  t2,
		t3:  t3,
		err: err,
	}
}

type mustWith3[T1, T2, T3 any] struct {
	t1  T1
	t2  T2
	t3  T3
	err error
}

func (m mustWith3[T1, T2, T3]) Replace(err error) (T1, T2, T3) {
	if m.err != nil {
		panic(err)
	}
	return m.t1, m.t2, m.t3
}

func (m mustWith3[T1, T2, T3]) Wrap(err error) (T1, T2, T3) {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
	return m.t1, m.t2, m.t3
}

func (m mustWith3[T1, T2, T3]) Callback(fn FnErrorCallback) (T1, T2, T3) {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
	return m.t1, m.t2, m.t3
}
