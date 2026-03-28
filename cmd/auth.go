package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/sahshad/apix/internal/config"
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
		fmt.Print("Enter API token: ")

		var token string
		fmt.Scanln(&token)
		if token == "" {
			cli.Error("Token cannot be empty.")
			return
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		cfg.AuthToken = token

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Error saving token:", err)
			return
		}

		cli.Success("Token saved successfully.")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove saved API token",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		if cfg.AuthToken == "" {
			cli.Warning("No token is currently set.")
			return
		}

		cfg.AuthToken = ""

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to remove token:", err)
			return
		}
		cli.Success("Token removed.")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current auth token status",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		if cfg.AuthToken != "" {
			cli.Success("Token is set.")
		} else {
			cli.Warning("No token set.")
		}
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(statusCmd)
}
