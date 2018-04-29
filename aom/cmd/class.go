package cmd

import (
	"github.com/spf13/cobra"
)

var module string

// classCmd represents the class command
var classCmd = &cobra.Command{
	Use:   "class",
	Short: "Manages the \"MetaModel\" resource",
	Long: `
class is a command for managing the \"MetaModel\" resource.
	
It's typically used during design-time of a module by a customer.

For example, to list all classes of module "MyModule":
aom class ls --module "MyModule" --baseUrl "https://apiomat.yourcompany.com/yambas/rest" --username "john" --password "secret"`,
}

func init() {
	rootCmd.AddCommand(classCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	classCmd.PersistentFlags().StringVar(&module, "module", "Basics", "Module name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// classCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
