package cmd

import (
	"fmt"

	"github.com/gobs/pretty"
	"github.com/spf13/cobra"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aoms"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists classes of a given module",
	Long: `
ls is a command for listing classes of a given module.

For example, to list all classes of module "MyModule":
aom class ls --module "MyModule" --baseUrl "https://apiomat.yourcompany.com/yambas/rest" --username "john" --password "secret"`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aomc.NewDefaultClient(baseUrl, username, password, aoms.System(system))

		classes, err := client.GetClasses(module)
		if err != nil {
			logFatal(err, "An error occurred during fetching the classes of module %v", module)
		}

		fmt.Printf("Classes of module %v:\n", module)
		pretty.PrettyPrint(classes)
	},
}

func init() {
	classCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
