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

var fnTest = func(inputError error) (int, error) { return 123, inputError }

func Example_Must(t *testing.T) {
	func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.Must(fnTest(nil))

		fmt.Printf("Output: %d\n", outputValue)

		return nil
	}()
	// Output: 123
}
