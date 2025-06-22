package tools_test

import (
	"errors"
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestRetry_SuccessAfterRetries(t *testing.T) {
	attempts := 5
	callCount := 0

	err := tools.Retry(attempts, func() error {
		callCount++
		if callCount < 3 {
			return errors.New("fail")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if callCount != 3 {
		t.Errorf("expected 3 calls, got %d", callCount)
	}
}

func TestRetry_FailsAllAttempts(t *testing.T) {
	attempts := 4
	callCount := 0
	expectedErr := errors.New("fail")

	err := tools.Retry(attempts, func() error {
		callCount++
		return expectedErr
	})

	if err != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
	if callCount != attempts {
		t.Errorf("expected %d calls, got %d", attempts, callCount)
	}
}

func TestRetry_ZeroAttempts(t *testing.T) {
	err := tools.Retry(0, func() error {
		t.Fatal("function should not be called")
		return nil
	})
	if err != nil {
		t.Errorf("expected nil error for zero attempts, got %v", err)
	}
}