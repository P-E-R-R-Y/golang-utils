package tools_test

import (
	"errors"
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestMust(t *testing.T) {
	want := 42
	got := tools.Must(want, nil)
	if got != want {
		t.Errorf("Must(nil error) = %v; want %v", got, want)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must(panic expected) did not panic")
		}
	}()
	tools.Must(0, errors.New("fail"))
}

func TestMustDo(t *testing.T) {
	tools.MustDo(nil) // should not panic

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustDo(panic expected) did not panic")
		}
	}()
	tools.MustDo(errors.New("fail"))
}

func TestSafe(t *testing.T) {
	val := 123
	got := tools.Safe(val, nil)
	if got != val {
		t.Errorf("Safe(nil error) = %v; want %v", got, val)
	}

	got = tools.Safe(val, errors.New("fail"))
	var zero int
	if got != zero {
		t.Errorf("Safe(non-nil error) = %v; want zero value %v", got, zero)
	}
}

func TestIgnore(t *testing.T) {
	// just test that Ignore does not panic
	tools.Ignore(errors.New("ignored error"))
	tools.Ignore(nil)
}

func TestWrapErr(t *testing.T) {
	err := errors.New("inner error")
	wrapped := tools.WrapErr("msg", err)
	if wrapped == nil {
		t.Errorf("WrapErr should not return nil for non-nil error")
	}
	if wrapped.Error() != "msg: inner error" {
		t.Errorf("WrapErr returned unexpected error string: %s", wrapped.Error())
	}

	if tools.WrapErr("msg", nil) != nil {
		t.Errorf("WrapErr should return nil when input error is nil")
	}
}

func TestRecover(t *testing.T) {
	// This just calls Recover inside a deferred function to verify it doesn't panic.
	defer tools.Recover()
	panic("test panic")
}