package yeh_test

import (
	"errors"
	"fmt"
	"github.com/hasanozgan/yeh"
	"strings"
	"testing"
)

func Test_MustWith_Replace_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1 := yeh.MustWith(fnTest(ErrFileNotFound)).Replace(ErrCustomFile)

		fmt.Printf("Output: %d\n", val1)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if got == ErrFileNotFound {
		t.Errorf("Failed, unexpected error")
	}

	if got != ErrCustomFile {
		t.Errorf("Failed, unexpected error")
	}
}

func Test_MustWith_Wrap_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1 := yeh.MustWith(fnTest(ErrFileNotFound)).Wrap(ErrCustomFile)

		fmt.Printf("Output: %d\n", val1)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if !strings.HasPrefix(got.Error(), ErrFileNotFound.Error()) {
		t.Errorf("Failed, unexpected error")
	}

	if !errors.Is(got, ErrCustomFile) {
		t.Errorf("Failed, unexpected error")
	}
}

func Test_MustWith_Callback_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1 := yeh.MustWith(fnTest(ErrFileNotFound)).Callback(func(err error) error {
			if errors.Is(err, ErrFileNotFound) {
				return ErrCustomFile
			}
			return err
		})

		fmt.Printf("Output: %d\n", val1)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if got == ErrFileNotFound {
		t.Errorf("Failed, unexpected error")
	}

	if got != ErrCustomFile {
		t.Errorf("Failed, unexpected error")
	}
}

func Test_MustWith_Replace_ShouldReturnValue(t *testing.T) {
	val1 := yeh.MustWith(fnTest(nil)).Replace(ErrCustomFile)

	if val1 != 123 {
		t.Error("Failed")
	}
}

func Test_MustWith_Wrap_ShouldReturnValue(t *testing.T) {
	val1 := yeh.MustWith(fnTest(nil)).Wrap(ErrCustomFile)

	if val1 != 123 {
		t.Error("Failed")
	}
}

func Test_MustWith_Callback_ShouldReturnValue(t *testing.T) {
	val1 := yeh.MustWith(fnTest(nil)).Callback(func(err error) error {
		if errors.Is(err, ErrFileNotFound) {
			return ErrCustomFile
		}
		return err
	})

	if val1 != 123 {
		t.Error("Failed")
	}
}
