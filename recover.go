package yeh

import (
	"errors"
	"fmt"
)

func Recover(err *error) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			appendError(err, fmt.Errorf("%s", r))
		case error:
			appendError(err, x)
		default:
			appendError(err, fmt.Errorf("%s", errors.New("unknown panic")))
		}
	}
}

func appendError(parent *error, err error) {
	if *parent == nil {
		*parent = err
	} else {
		*parent = fmt.Errorf("%s: %w", (*parent).Error(), err)
	}
}
