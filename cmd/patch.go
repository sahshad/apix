package cmd

import (
	"fmt"
	"os"

	"github.com/sahshad/apix/internal/cli"
	"github.com/spf13/cobra"
)

type PatchCmdOptions struct {
	Data      string
	File      string
	FormFiles []string
}

var patchOpts PatchCmdOptions

var patchCmd = &cobra.Command{
	Use:   "patch [endpoint]",
	Short: "PATCH data to API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		method := "PATCH"
		endpoint := args[0]
		client := cli.GetClient()

		if len(patchOpts.FormFiles) > 0 {
			res, err := client.Multipart(method, endpoint, postOpts.FormFiles)
			if err != nil {
				cli.Error("Request failed:", err)
				return
			}

			resParams := cli.BuildResponseParams(method, endpoint, res)
			cli.RenderResponse(resParams, verbose)
			return
		}

		payload := patchOpts.Data
		if patchOpts.File != "" {
			content, err := os.ReadFile(patchOpts.File)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			payload = string(content)
		}

		res, err := client.Patch(endpoint, payload)
		if err != nil {
			cli.Error("Request failed:", err)
			return
		}

		resParams := cli.BuildResponseParams(method, endpoint, res)
		cli.RenderResponse(resParams, verbose)
	},
}

func init() {
	patchCmd.Flags().StringVarP(&patchOpts.Data, "data", "d", "", "JSON payload")
	patchCmd.Flags().StringVarP(&patchOpts.File, "file", "f", "", "Payload from file")
	patchCmd.Flags().StringArrayVarP(&patchOpts.FormFiles, "form-file", "F", []string{}, "Multipart form file (format: field=path/to/file)")
}
