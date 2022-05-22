package yeh_test

import (
	"fmt"
	"github.com/hasanozgan/yeh"

	"testing"
)

func Test_Must0_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		yeh.Must0(fnTest0(ErrFileNotFound))

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value := yeh.Must(fnTest(ErrFileNotFound))

		fmt.Printf("Output: %d\n", value)

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must2_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value1, value2 := yeh.Must2(fnTest2(ErrFileNotFound))

		fmt.Printf("Output: %d %d\n", value1, value2)

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must3_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value1, value2, value3 := yeh.Must3(fnTest3(ErrFileNotFound))

		fmt.Printf("Output: %d %d %d\n", value1, value2, value3)

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must4_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value1, value2, value3, value4 := yeh.Must4(fnTest4(ErrFileNotFound))

		fmt.Printf("Output: %d %d %d %d\n", value1, value2, value3, value4)

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must5_ShouldReturnError(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value1, value2, value3, value4, value5 := yeh.Must5(fnTest5(ErrFileNotFound))

		fmt.Printf("Output: %d %d %d %d %d\n", value1, value2, value3, value4, value5)

		return nil
	}()

	if got == nil {
		t.Error("Failed")
	}

	if got != ErrFileNotFound {
		t.Error("Failed")
	}
}

func Test_Must_ShouldReturnValue(t *testing.T) {
	value1 := yeh.Must(fnTest(nil))

	if value1 != 123 {
		t.Error("Failed")
	}
}

func Test_Must2_ShouldReturnValue(t *testing.T) {
	value1, value2 := yeh.Must2(fnTest2(nil))

	if value1 != 123 {
		t.Error("Failed")
	}

	if value2 != 234 {
		t.Error("Failed")
	}
}

func Test_Must3_ShouldReturnValue(t *testing.T) {
	value1, value2, value3 := yeh.Must3(fnTest3(nil))

	if value1 != 123 {
		t.Error("Failed")
	}

	if value2 != 234 {
		t.Error("Failed")
	}

	if value3 != 345 {
		t.Error("Failed")
	}
}

func Test_Must4_ShouldReturnValue(t *testing.T) {
	value1, value2, value3, value4 := yeh.Must4(fnTest4(nil))

	if value1 != 123 {
		t.Error("Failed")
	}

	if value2 != 234 {
		t.Error("Failed")
	}

	if value3 != 345 {
		t.Error("Failed")
	}

	if value4 != 456 {
		t.Error("Failed")
	}
}

func Test_Must5_ShouldReturnValue(t *testing.T) {
	value1, value2, value3, value4, value5 := yeh.Must5(fnTest5(nil))

	if value1 != 123 {
		t.Error("Failed")
	}

	if value2 != 234 {
		t.Error("Failed")
	}

	if value3 != 345 {
		t.Error("Failed")
	}

	if value4 != 456 {
		t.Error("Failed")
	}

	if value5 != 567 {
		t.Error("Failed")
	}
}
