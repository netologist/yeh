package yeh_test

import (
	"errors"
	"fmt"
	"github.com/hasanozgan/yeh"
	"strings"
	"testing"
)

func Test_MustWith5_Replace_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(ErrFileNotFound)).Replace(ErrCustomFile)

		fmt.Printf("Output: %d %d %d %d %d\n", val1, val2, val3, val4, val5)

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

func Test_MustWith5_Wrap_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(ErrFileNotFound)).Wrap(ErrCustomFile)

		fmt.Printf("Output: %d %d %d %d %d\n", val1, val2, val3, val4, val5)

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

func Test_MustWith5_Callback_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(ErrFileNotFound)).Callback(func(err error) error {
			if errors.Is(err, ErrFileNotFound) {
				return ErrCustomFile
			}
			return err
		})

		fmt.Printf("Output: %d %d %d %d %d\n", val1, val2, val3, val4, val5)

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

func Test_MustWith5_Replace_ShouldReturnValue(t *testing.T) {
	val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(nil)).Replace(ErrCustomFile)

	if val1 != 123 {
		t.Error("Failed")
	}

	if val2 != 234 {
		t.Error("Failed")
	}

	if val3 != 345 {
		t.Error("Failed")
	}

	if val4 != 456 {
		t.Error("Failed")
	}

	if val5 != 567 {
		t.Error("Failed")
	}
}

func Test_MustWith5_Wrap_ShouldReturnValue(t *testing.T) {
	val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(nil)).Wrap(ErrCustomFile)

	if val1 != 123 {
		t.Error("Failed")
	}

	if val2 != 234 {
		t.Error("Failed")
	}

	if val3 != 345 {
		t.Error("Failed")
	}

	if val4 != 456 {
		t.Error("Failed")
	}

	if val5 != 567 {
		t.Error("Failed")
	}
}

func Test_MustWith5_Callback_ShouldReturnValue(t *testing.T) {
	val1, val2, val3, val4, val5 := yeh.MustWith5(fnTest5(nil)).Callback(func(err error) error {
		if errors.Is(err, ErrFileNotFound) {
			return ErrCustomFile
		}
		return err
	})

	if val1 != 123 {
		t.Error("Failed")
	}

	if val2 != 234 {
		t.Error("Failed")
	}

	if val3 != 345 {
		t.Error("Failed")
	}

	if val4 != 456 {
		t.Error("Failed")
	}

	if val5 != 567 {
		t.Error("Failed")
	}
}
