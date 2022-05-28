# yeh
**Y**et another **E**rror **H**andler for Golang. 

*This Project An **Experimental** Idea For Error Handler Implementation With Go Generics*

## Examples

### Must

Also Must has support *Must0, Must2, Must3, Must4* and *Must5* methods. For multiple return types.

```go

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

```

### MustWith

Also Must has support *MustWith0, MustWith2, MustWith3, MustWith4* and *MustWith5* methods. For multiple return types.

#### Replace
```go

got := func() (err error) {
    defer yeh.Recover(&err)

    outputValue := yeh.MustWith(os.Open("file.not.exists.txt"))
    .Replace(ErrCustomExists)

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

```

#### Wrap
```go

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

```

#### Callback
```go

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

```
