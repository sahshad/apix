package cmd

import (
	"apix/internal/client"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var data string
var file string

var postCmd = &cobra.Command{
    Use:   "post [endpoint]",
    Short: "POST data to API",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        endpoint := args[0]

        payload := data
        if file != "" {
            content, err := os.ReadFile(file)
            if err != nil {
                fmt.Println("Error reading file:", err)
                return
            }
            payload = string(content)
        }

        resp, err := client.Post(endpoint, payload)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        fmt.Println(string(resp))
    },
}

func init() {
    postCmd.Flags().StringVarP(&data, "data", "d", "", "JSON payload")
    postCmd.Flags().StringVarP(&file, "file", "f", "", "Payload from file")
}