package yeh

import "fmt"

func MustWith0(err error) mustWith0 {
	return mustWith0{
		err: err,
	}
}

type mustWith0 struct {
	err error
}

func (m mustWith0) Replace(err error) {
	if m.err != nil {
		panic(err)
	}
}

func (m mustWith0) Wrap(err error) {
	if m.err != nil {
		panic(fmt.Errorf("%s: [%w]", m.err.Error(), err))
	}
}

func (m mustWith0) Callback(fn FnErrorCallback) {
	if cerr := fn(m.err); cerr != nil {
		panic(cerr)
	}
}
