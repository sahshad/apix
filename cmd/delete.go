package cmd

import (
	"apix/internal/client"
	"fmt"

	"github.com/spf13/cobra"
)

var force bool

var deleteCmd = &cobra.Command{
    Use:   "delete [endpoint]",
    Short: "DELETE resource from API",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        endpoint := args[0]
        if !force {
            var confirm string
            fmt.Printf("Are you sure you want to delete %s? (y/N): ", endpoint)
            fmt.Scanln(&confirm)
            if confirm != "y" && confirm != "Y" {
                fmt.Println("Aborted.")
                return
            }
        }

        resp, err := client.Delete(endpoint)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(string(resp))
    },
}

func init() {
    deleteCmd.Flags().BoolVarP(&force, "force", "f", false, "Force deletion without confirmation")
}