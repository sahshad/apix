package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/spf13/cobra"
)

type GetCmdOptions struct {
	Output string
	Pretty bool
}

var getOpts GetCmdOptions

var getCmd = &cobra.Command{
	Use:   "get [endpoint]",
	Short: "GET data from API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		method := "GET"
		endpoint := args[0]
		c := cli.GetClient()

		res, err := c.Get(endpoint)
		if err != nil {
			cli.Error("Request failed:", err)
			return
		}

		resParams := cli.BuildResponseParams(method, endpoint, res)
		cli.RenderResponse(resParams, verbose)

		// save file
		if getOpts.Output != "" {
			err := cli.SaveToFile(getOpts.Output, res.Body)
			if err != nil {
				cli.Error("Failed to save file:", err)
			} else {
				cli.Success("Saved response to", getOpts.Output)
			}
		}
	},
}

func init() {
	getCmd.Flags().StringVarP(&getOpts.Output, "output", "o", "", "Save response to file")
	getCmd.Flags().BoolVarP(&getOpts.Pretty, "pretty", "p", false, "Pretty print JSON")
}
