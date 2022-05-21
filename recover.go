package yeh

import (
	"errors"
	"fmt"
)

func Recover(err *error) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			if *err == nil {
				*err = fmt.Errorf("%s", r)
			} else {
				*err = fmt.Errorf("%s: %s", (*err).Error(), r)
			}
		case error:
			if *err == nil {
				*err = x
			} else {
				*err = fmt.Errorf("%s: %w", (*err).Error(), x)
			}
		default:
			*err = errors.New("unknown panic")
		}
	}
}
