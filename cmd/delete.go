package cmd

import (
	"apix/internal/cli"
	"apix/internal/types"
	"fmt"
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
		endpoint := args[0]

		if !deleteOpts.Force {
			var confirm string
			fmt.Printf("Are you sure you want to delete '%s'? (y/N): ", endpoint)
			fmt.Scanln(&confirm)

			if confirm != "y" && confirm != "Y" {
				cli.Warning("Aborted.")
				return
			}
		}

		c := cli.GetClient()

		res, err := c.Delete(endpoint)
		if err != nil {
			cli.Error("Request failed:", err)
			return
		}

		responseParams := types.ResponseParams{
			Method:      "GET",
			Endpoint:    endpoint,
			Status:      res.StatusCode,
			ContentType: res.Headers.Get("Content-Type"),
			Body:        string(res.Body),
			Duration:    res.DurationMs,
			Size:        cli.FormatSize(res.Size),
		}

		cli.RenderResponse(responseParams, verbose)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&deleteOpts.Force, "force", "f", false, "Force deletion without confirmation")
}
