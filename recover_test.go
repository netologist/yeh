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
		t.Error("Failed")
	}

	if !errors.Is(got, ErrFileNotFound) {
		t.Error("Failed")
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
		t.Error("Failed")
	}

	if !strings.HasPrefix(got.Error(), ErrDBConnection.Error()) {
		t.Error("Failed, unexpected error")
	}

	if !errors.Is(got, ErrFileNotFound) {
		t.Error("Failed")
	}
}

func Test_Recover_String(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		panic("some error string")

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if !strings.ContainsAny(got.Error(), "some error string") {

		t.Error("Failed")
	}
}

func Test_Recover_Unknown(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		panic(struct{}{})

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if !strings.ContainsAny(got.Error(), "unknown panic") {

		t.Error("Failed")
	}
}
