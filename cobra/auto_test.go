package cobra_test

import (
	"testing"

	"github.com/spf13/cobra"
	mycobra "github.com/P-E-R-R-Y/golang-utils/cobra"
)

func TestGetAutoCompletionAddCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "testcli"}

	// Add to root command
	rootCmd.AddCommand(mycobra.GetAutoCompletion("testtool"))

	// Check if it was added successfully
	found := false
	for _, c := range rootCmd.Commands() {
		if c.Use == "auto" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected 'auto' command to be added to root")
	}
}


func TestGetManualCompletionAddCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "testcli"}

	// Add to root command
	rootCmd.AddCommand(mycobra.GetManualCompletion())

	// Check if it was added successfully
	found := false
	for _, c := range rootCmd.Commands() {
		if c.Use == "manual" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected 'manual' command to be added to root")
	}
}