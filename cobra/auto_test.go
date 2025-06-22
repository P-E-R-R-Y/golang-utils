package cobra_test

import (
	"testing"

	"github.com/spf13/cobra"
	mycobra "github.com/P-E-R-R-Y/golang-utils/cobra"
)

func TestGetAutoCompletionAddCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "testcli"}

	// Get the auto-completion command with a custom name
	autoCmd := mycobra.GetAutoCompletion("testtool")

	// Add to root command
	rootCmd.AddCommand(autoCmd)

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

	// Get the auto-completion command with a custom name
	manualCmd := mycobra.GetManualCompletion()

	// Add to root command
	rootCmd.AddCommand(manualCmd)

	// Check if it was added successfully
	found := false
	for _, c := range rootCmd.Commands() {
		if c.Use == "manual" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected 'auto' command to be added to root")
	}
}