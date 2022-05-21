package yeh_test

import (
	"errors"
	"fmt"
	"github.com/hasanozgan/yeh"
	"strings"
	"testing"
)

func Test_Must(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value := yeh.Must(fnTest(ErrFileNotFound))

		fmt.Printf("Output: %d\n", value)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed")
	}

	if got != ErrFileNotFound {
		t.Errorf("Failed")
	}
}

func Test_MustWith_Replace(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(fnTest(ErrFileNotFound)).Replace(ErrCustomFile)

		fmt.Printf("Output: %d\n", outputValue)

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

func Test_MustWith_Wrap(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(fnTest(ErrFileNotFound)).Wrap(ErrCustomFile)

		fmt.Printf("Output: %d\n", outputValue)

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

func Test_MustWith_Callback(t *testing.T) {

	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(fnTest(ErrFileNotFound)).Callback(func(err error) error {
			if errors.Is(err, ErrFileNotFound) {
				return ErrCustomFile
			}
			return err
		})

		fmt.Printf("Output: %d\n", outputValue)

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
