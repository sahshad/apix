package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "apix",
    Short: "apix - A better curl for APIs",
    Long:  "apix is a CLI tool to interact with APIs easily, manage environments, auth, and output.",
}

var verbose bool
var env string

func Execute() {
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
    rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "default", "Select environment")

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(getCmd)
    rootCmd.AddCommand(postCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(authCmd)
    rootCmd.AddCommand(envCmd)
    rootCmd.AddCommand(configCmd)
}