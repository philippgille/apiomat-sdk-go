package cmd

import (
	"fmt"

	"github.com/gobs/pretty"
	"github.com/spf13/cobra"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aomx"
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
		client := aomc.NewDefaultClient(baseUrl, username, password, aomx.System(system))

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
}
