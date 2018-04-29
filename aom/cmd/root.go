package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var baseUrl string
var username string
var password string
var system string

var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aom",
	Short: "A CLI for managing ApiOmat customer and user resources",
	Long: `
aom is a CLI for managing ApiOmat resources.

This includes design-time resources for customers,
as well as runtime resources for users.

Examples:

- Print the version:
	aom version --baseUrl "https://apiomat.yourcompany.com/yambas/rest"
- List all classes of module "MyModule":
	aom class ls --module "MyModule" --baseUrl "https://apiomat.yourcompany.com/yambas/rest" --username "john" --password "secret"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&baseUrl, "baseUrl", "http://localhost:8080/yambas/rest", "Base URL")
	rootCmd.PersistentFlags().StringVar(&username, "username", "apinaut", "Username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "secret", "Password")
	rootCmd.PersistentFlags().StringVar(&system, "system", "", "System (no default value, leads to the ApiOmat server using LIVE)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
