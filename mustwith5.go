package yeh

import "fmt"

func MustWith5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) mustWith5[T1, T2, T3, T4, T5] {
	return mustWith5[T1, T2, T3, T4, T5]{
		t1:  t1,
		t2:  t2,
		t3:  t3,
		t4:  t4,
		t5:  t5,
		err: err,
	}
}

type mustWith5[T1, T2, T3, T4, T5 any] struct {
	t1  T1
	t2  T2
	t3  T3
	t4  T4
	t5  T5
	err error
}

func (m mustWith5[T1, T2, T3, T4, T5]) Replace(err error) (T1, T2, T3, T4, T5) {
	if m.err != nil {
		panic(err)
	}
	return m.t1, m.t2, m.t3, m.t4, m.t5
}

func (m mustWith5[T1, T2, T3, T4, T5]) Wrap(err error) (T1, T2, T3, T4, T5) {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
	return m.t1, m.t2, m.t3, m.t4, m.t5
}

func (m mustWith5[T1, T2, T3, T4, T5]) Callback(fn FnErrorCallback) (T1, T2, T3, T4, T5) {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
	return m.t1, m.t2, m.t3, m.t4, m.t5
}
