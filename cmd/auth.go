package cmd

import (
	"apix/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication tokens",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login and save API token",
	Run: func(cmd *cobra.Command, args []string) {
		var token string
		fmt.Print("Enter API token: ")
		fmt.Scanln(&token)

		cfg, _ := config.LoadConfig()
		cfg.AuthToken = token
		err := config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("Error saving token:", err)
			return
		}
		fmt.Println("Token saved successfully.")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove saved API token",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		cfg.AuthToken = ""
		config.SaveConfig(cfg)
		fmt.Println("Token removed.")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current auth token status",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		if cfg.AuthToken != "" {
			fmt.Println("Token is set.")
		} else {
			fmt.Println("No token set.")
		}
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(statusCmd)
}