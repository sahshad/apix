package cmd

import (
	"apix/internal/client"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var filter string
var limit int
var output string
var pretty bool

var getCmd = &cobra.Command{
    Use:   "get [endpoint]",
    Short: "GET data from API",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        endpoint := args[0]
        resp, err := client.Get(endpoint)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        if pretty {
            var prettyJSON map[string]interface{}
            json.Unmarshal(resp, &prettyJSON)
            formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
            fmt.Println(string(formatted))
        } else {
            fmt.Println(string(resp))
        }

        if output != "" {
            client.SaveToFile(output, resp)
        }
    },
}

func init() {
    getCmd.Flags().StringVar(&filter, "filter", "", "Filter results")
    getCmd.Flags().IntVar(&limit, "limit", 0, "Limit number of results")
    getCmd.Flags().StringVarP(&output, "output", "o", "", "Save response to file")
    getCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print JSON")
}