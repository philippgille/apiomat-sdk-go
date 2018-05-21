package cmd

import (
	"github.com/spf13/cobra"
)

var module string

// classCmd represents the class command
var classCmd = &cobra.Command{
	Use:   "class",
	Short: `Manages the "MetaModel" resource`,
	Long: `
class is a command for managing the "MetaModel" resource.
	
It's typically used during design-time of a module by a customer.

For example, to list all classes of module "MyModule":
aom class ls --module "MyModule" --baseUrl "https://apiomat.yourcompany.com/yambas/rest" --username "john" --password "secret"`,
}

func init() {
	rootCmd.AddCommand(classCmd)

	// Persistent flags

	registerCommonFlags(classCmd)

	classCmd.PersistentFlags().StringVar(&module, "module", "Basics", "Module name")
}
