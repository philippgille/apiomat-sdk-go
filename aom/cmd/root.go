package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Increment and remove "+" in release commits.
// Add "+" after release commits.
const version = "v0.3.0+"

var baseUrl string
var username string
var password string
var system string

var debug bool
var printVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aom",
	Short: "A CLI for managing ApiOmat customer and user resources",
	Long: `
aom is a CLI for managing ApiOmat resources.

This includes design-time resources for customers,
as well as runtime resources for users.

Examples:

- Print the version of the ApiOmat server:
	aom version --baseUrl "https://apiomat.yourcompany.com/yambas/rest"
- List all classes of module "MyModule":
	aom class ls --module "MyModule" --baseUrl "https://apiomat.yourcompany.com/yambas/rest" --username "john" --password "secret"`,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Printf("aom version: %v", version)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Persistent flags

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug switch. Activate to include stack trace when errors are logged")

	// Local flags

	// Use our own version flag instead of setting "Version" on the rootCmd to have a nicer message and the "-v" shorthand
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Print the version of the aom CLI (not of the ApiOmat server - use \"aom version\" for that")
}

// logError uses the CLI's debug flag to either print the error with stack trace or without.
// Calling this:
// logError(err, "An error occurred during fetching the classes of module %v", module )
// Leads to this:
// log.Printf("An error occurred during fetching the classes of module %v: %v", module, err)
// Or if the debug flag is activated, the error gets formatted with %+v, leading to the stack trace being logged as well.
func logError(err error, format string, v ...interface{}) {
	// Even if no third parameter was passed, v is an empty slice, so appending works
	v = append(v, err)
	if debug {
		log.Printf(format+":\n%+v", v...)
	} else {
		log.Printf(format+":\n%v", v...)
	}
}

// logFatal executes logError and then os.Exit(1)
func logFatal(err error, format string, v ...interface{}) {
	logError(err, format, v...)
	os.Exit(1)
}

func registerBaseUrl(command *cobra.Command) {
	command.PersistentFlags().StringVar(&baseUrl, "baseUrl", "http://localhost:8080/yambas/rest", "Base URL")
}

func registerCommonFlags(command *cobra.Command) {
	registerBaseUrl(command)
	command.PersistentFlags().StringVar(&username, "username", "apinaut", "Username")
	command.PersistentFlags().StringVar(&password, "password", "secret", "Password")
	command.PersistentFlags().StringVar(&system, "system", "", "System (no default value, leads to the ApiOmat server using LIVE)")
}
