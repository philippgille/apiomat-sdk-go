package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aoms"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the ApiOmat server",
	Long: `version is a command for printing the version of the ApiOmat server.

For example:
aom version --baseUrl "https://apiomat.yourcompany.com/yambas/rest"`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aomc.NewDefaultClient(baseUrl, username, password, aoms.System(system))

		version, err := client.GetVersion()
		if err != nil {
			panic(err)
		}

		fmt.Printf("ApiOmat version: %v", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
