package yeh

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
