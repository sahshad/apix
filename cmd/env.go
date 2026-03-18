package cmd

import (
	"apix/internal/config"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage API environments",
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available environments",
	Run: func(cmd *cobra.Command, args []string) {
		// For simplicity, hardcoded example
		fmt.Println("Available environments:")
		fmt.Println("- default (https://api.example.com)")
		fmt.Println("- staging (https://staging.api.example.com)")
		fmt.Println("- dev (https://dev.api.example.com)")
	},
}

var envUseCmd = &cobra.Command{
	Use:   "use [environment]",
	Short: "Switch to a specific environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName := args[0]
		cfg, _ := config.LoadConfig()

		switch envName {
		case "default":
			cfg.BaseURL = "https://api.example.com"
		case "staging":
			cfg.BaseURL = "https://staging.api.example.com"
		case "dev":
			cfg.BaseURL = "https://dev.api.example.com"
		default:
			fmt.Println("Unknown environment:", envName)
			return
		}

		config.SaveConfig(cfg)
		fmt.Printf("Switched to environment: %s\n", envName)
	},
}

var envSetCmd = &cobra.Command{
	Use:   "set [KEY=VALUE]",
	Short: "Set environment variable",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kv := args[0]
		cfg, _ := config.LoadConfig()

		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			color.Red("Invalid format. Use KEY=VALUE")
			return
		}
		key := parts[0]
		value := parts[1]

		switch key {
		case "API_URL":
			cfg.BaseURL = value
		default:
			fmt.Println("Unknown key:", key)
			return
		}
		config.SaveConfig(cfg)
		fmt.Println("Environment variable set:", kv)
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envUseCmd)
	envCmd.AddCommand(envSetCmd)
}
