package yeh_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hasanozgan/yeh"
)

var ErrFileNotFound = errors.New("file not found")
var ErrCustomFile = errors.New("custom file")
var ErrDBConnection = errors.New("db connection")

var fnTest0 = func(inputError error) error { return inputError }
var fnTest = func(inputError error) (int, error) { return 123, inputError }
var fnTest2 = func(inputError error) (int, int, error) { return 123, 234, inputError }
var fnTest3 = func(inputError error) (int, int, int, error) { return 123, 234, 345, inputError }
var fnTest4 = func(inputError error) (int, int, int, int, error) { return 123, 234, 345, 456, inputError }
var fnTest5 = func(inputError error) (int, int, int, int, int, error) { return 123, 234, 345, 456, 567, inputError }

func Example_Must(t *testing.T) {
	func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.Must(fnTest(nil))

		fmt.Printf("Output: %d\n", outputValue)

		return nil
	}()
	// Output: 123
}
