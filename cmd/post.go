package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/sahshad/apix/internal/types"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type PostCmdOptions struct {
	Data   string
	File   string
	Pretty bool
}

var postOpts PostCmdOptions

var postCmd = &cobra.Command{
	Use:   "post [endpoint]",
	Short: "POST data to API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := args[0]

		payload := postOpts.Data
		if postOpts.File != "" {
			content, err := os.ReadFile(postOpts.File)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			payload = string(content)
		}

		client := cli.GetClient()

		res, err := client.Post(endpoint, payload)
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
			Timing:      res.Timing,
		}

		cli.RenderResponse(responseParams, verbose)

	},
}

func init() {
	postCmd.Flags().StringVarP(&postOpts.Data, "data", "d", "", "JSON payload")
	postCmd.Flags().StringVarP(&postOpts.File, "file", "f", "", "Payload from file")
}
