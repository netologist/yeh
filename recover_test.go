package yeh_test

import (
	"errors"
	"fmt"
	"github.com/hasanozgan/yeh"
	"strings"
	"testing"
)

func Test_Recover(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value := yeh.Must(fnTest(ErrFileNotFound))

		fmt.Printf("Output: %d\n", value)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed")
	}

	if !errors.Is(got, ErrFileNotFound) {
		t.Errorf("Failed")
	}
}

func Test_Recover_Wrap(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		err = ErrDBConnection

		value := yeh.Must(fnTest(ErrFileNotFound))

		fmt.Printf("Output: %d\n", value)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed")
	}

	if !strings.HasPrefix(got.Error(), ErrDBConnection.Error()) {
		t.Errorf("Failed, unexpected error")
	}

	if !errors.Is(got, ErrFileNotFound) {
		t.Errorf("Failed")
	}
}
