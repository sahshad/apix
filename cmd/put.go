package cmd

import (
	"fmt"
	"os"

	"github.com/sahshad/apix/internal/cli"
	"github.com/spf13/cobra"
)

type PutCmdOptions struct {
	Data      string
	File      string
	FormFiles []string
}

var putOpts PutCmdOptions

var putCmd = &cobra.Command{
	Use:   "put [endpoint]",
	Short: "PUT Data to API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		method := "PUT"
		endpoint := args[0]
		client := cli.GetClient()

		if len(putOpts.FormFiles) > 0 {
			res, err := client.Multipart(method, endpoint, postOpts.FormFiles)
			if err != nil {
				cli.Error("Request failed:", err)
				return
			}

			resParams := cli.BuildResponseParams(method, endpoint, res)
			cli.RenderResponse(resParams, verbose)
			return
		}

		payload := putOpts.Data
		if putOpts.File != "" {
			content, err := os.ReadFile(putOpts.File)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			payload = string(content)
		}

		res, err := client.Put(endpoint, payload)
		if err != nil {
			cli.Error("Request failed:", err)
			return
		}

		resParams := cli.BuildResponseParams(method, endpoint, res)
		cli.RenderResponse(resParams, verbose)
	},
}

func init() {
	putCmd.Flags().StringVarP(&putOpts.Data, "data", "d", "", "JSON Payload")
	putCmd.Flags().StringVarP(&putOpts.File, "file", "f", "", "Payload from file")
	putCmd.Flags().StringArrayVarP(&putOpts.FormFiles, "form-file", "F", []string{}, "Multipart form file (format: field=path/to/file)")
}
