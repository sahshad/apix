package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/sahshad/apix/internal/types"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

type GetCmdOptions struct {
	Filter string
	Limit  int
	Output string
	Pretty bool
}

var getOpts GetCmdOptions

var getCmd = &cobra.Command{
	Use:   "get [endpoint]",
	Short: "GET data from API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := args[0]

		// build query params
		query := url.Values{}
		if getOpts.Filter != "" {
			query.Add("filter", getOpts.Filter)
		}
		if getOpts.Limit > 0 {
			query.Add("limit", fmt.Sprintf("%d", getOpts.Limit))
		}

		if len(query) > 0 {
			endpoint = endpoint + "?" + query.Encode()
		}

		c := cli.GetClient()

		res, err := c.Get(endpoint)
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
			Headers:     res.Headers,
			Timing:      res.Timing,
		}

		cli.RenderResponse(responseParams, verbose)

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
	getCmd.Flags().StringVar(&getOpts.Filter, "filter", "", "Filter results")
	getCmd.Flags().IntVar(&getOpts.Limit, "limit", 0, "Limit number of results")
	getCmd.Flags().StringVarP(&getOpts.Output, "output", "o", "", "Save response to file")
	getCmd.Flags().BoolVarP(&getOpts.Pretty, "pretty", "p", false, "Pretty print JSON")
}
