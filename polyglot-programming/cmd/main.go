package main

import "fmt"

func returnsError(value int) error {
	return fmt.Errorf("this is an error with val: %v", value)
}

type Foo struct { }

func (f *Foo) thisIsOnFoo() error {
	return fmt.Errorf("this is an error with val.")
}

func CreateFoo(fail bool) (*Foo, error) {
	if fail {
		return nil, fmt.Errorf("this is an error with val.")
	}

	return &Foo{}, nil
}

func main() {
	foo, err := CreateFoo(false)

	if err != nil {
		// ...
	}
}
