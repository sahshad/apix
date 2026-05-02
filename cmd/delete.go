package cmd

import (
	"fmt"
	"github.com/sahshad/apix/internal/cli"
	"github.com/spf13/cobra"
)

type DeleteCmdOptions struct {
	Force bool
}

var deleteOpts DeleteCmdOptions

var deleteCmd = &cobra.Command{
	Use:   "delete [endpoint]",
	Short: "DELETE resource from API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		method := "DELETE"
		endpoint := args[0]
		client := cli.GetClient()

		if !deleteOpts.Force {
			var confirm string
			fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", endpoint)
			fmt.Scanln(&confirm)

			if confirm != "y" && confirm != "Y" {
				cli.Warning("Aborted.")
				return
			}
		}

		res, err := client.Delete(endpoint)
		if err != nil {
			cli.Error("Request failed:", err)
			return
		}

		resParams := cli.BuildResponseParams(method, endpoint, res)
		cli.RenderResponse(resParams, verbose)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&deleteOpts.Force, "force", "f", false, "Force deletion without confirmation")
}
