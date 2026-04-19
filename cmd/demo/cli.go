package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// This will be set with -ldflags="-X main.version=value" at the build time.
// It need not be exported. Hence, keeping in lowercase.
var version string

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo is a sample application",
	Long:  "demo is a sample application that demonstrates production grade boilerplate.",
	// Often, we want to control how errors are logged/printed. Its usually a good idea to silence the default error printing.
	SilenceErrors: true,
	// Don't set SilenceUsage as "true" here. However, "Usage" will be displayed for any error encountered in the RunE function.
	// We should ideally display usage only for invalid User Options/Values/Args and not for other logic errors.
	// So, we can keep SilenceUsage as false and later in the RunE function, after all necessary validation, we can SilenceUsage as true.
	SilenceUsage: false,
	// There's a "Run" function as well but using RunE is recommended as it
	// allows you to return an error. This enables custom error handling and
	// proper exit codes, which is important for scripting and automation.
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 && args[0] == "simulate-validation-error" {
			return errors.New("simulated error that shows up with usage message")
		}
		// Turn off usage related errors after validation.
		cmd.SilenceUsage = true
		if len(args) >= 1 && args[0] == "simulate-logic-error" {
			return errors.New("simulated logic error that doesn't show usage")
		}
		fmt.Println("Version:", version)
		return nil
	},
}

// main function in main package.
// filename doesn't matter.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
