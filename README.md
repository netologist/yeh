# yeh
<span style="color: #db4c3f; font-weight: 700">Y</span>et another <span style="color: #db4c3f; font-weight: 700">E</span>rror <span style="color: #db4c3f; font-weight: 700">H</span>andler for Golang. 

*This Project An **Experimental** Idea For Error Handling Implementation With Go Generics*

## Examples

### Must

Also Must has support Must0, Must2, Must3, Must4 and Must5 methods. For multiple return types.

```go
func Test_Must(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		value := yeh.Must(os.Open("file.not.exists.txt"))

		fmt.Printf("Output: %d\n", value)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed")
	}

	if got != fs.ErrExist {
		t.Errorf("Failed")
	}
}
```

### MustWith.Replace
```go
func Test_MustWith_Replace(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(os.Open("file.not.exists.txt")).Replace(ErrCustomExists)

		fmt.Printf("Output: %d\n", outputValue)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if got == fs.ErrExist {
		t.Errorf("Failed, unexpected error")
	}

	if got != ErrCustomExists {
		t.Errorf("Failed, unexpected error")
	}
}
```

### MustWith.Wrap
```go
func Test_MustWith_Wrap(t *testing.T) {
	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(os.Open("file.not.exists.txt")).Wrap(ErrCustomExists)

		fmt.Printf("Output: %d\n", outputValue)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if !strings.HasPrefix(got.Error(), fs.ErrExist.Error()) {
		t.Errorf("Failed, unexpected error")
	}

	if !errors.Is(got, ErrCustomExists) {
		t.Errorf("Failed, unexpected error")
	}
}
```

### MustWith.Callback
```go
func Test_MustWith_Callback(t *testing.T) {

	got := func() (err error) {
		defer yeh.Recover(&err)

		outputValue := yeh.MustWith(os.Open("file.not.exists.txt")).Callback(func(err error) error {
			if errors.Is(err, fs.) {
				return ErrCustomExists
			}
			return err
		})

		fmt.Printf("Output: %d\n", outputValue)

		return nil
	}()

	if got == nil {
		t.Errorf("Failed, unexpected error")
	}

	if got == fs.ErrExist {
		t.Errorf("Failed, unexpected error")
	}

	if got != ErrCustomExists {
		t.Errorf("Failed, unexpected error")
	}
}
```
