package yeh_test

import (
	"errors"
	"github.com/hasanozgan/yeh"
	"strings"
	"testing"
)

func Test_MustWith0_Replace(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		yeh.MustWith0(fnTest0(ErrFileNotFound)).Replace(ErrCustomFile)

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

func Test_MustWith0_Wrap(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		yeh.MustWith0(fnTest0(ErrFileNotFound)).Wrap(ErrCustomFile)

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

func Test_MustWith0_Callback(t *testing.T) {

	got := func() (err error) {
		defer yeh.Recover(&err)

		yeh.MustWith0(fnTest0(ErrFileNotFound)).Callback(func(err error) error {
			if errors.Is(err, ErrFileNotFound) {
				return ErrCustomFile
			}
			return err
		})

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
