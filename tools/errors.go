package tools

import (
	"fmt"
)

// Must panics if err is not nil, else returns val.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func MustDo(err error) {
	if err != nil {
		panic(err)
	}
}

// Safe returns zero value of T if err is not nil, else returns val.
func Safe[T any](val T, err error) T {
	var zero T
	if err != nil {
		return zero
	}
	return val
}

func Ignore(err error) {
	_ = err
}

// WrapErr wraps err with msg if err is not nil.
func WrapErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}

//Recover
func Recover() {
	if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}