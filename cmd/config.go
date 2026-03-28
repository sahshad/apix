package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/sahshad/apix/internal/config"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current CLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		b, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			cli.Error("Failed to format config:", err)
			return
		}

		fmt.Println(string(b))
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set KEY=VALUE",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			cli.Error("Invalid format. Use KEY=VALUE")
			return
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "base_url":
			cfg.BaseURL = value

		case "auth_token":
			cfg.AuthToken = value

		default:
			cli.Warning("Unknown config key:", key)
			return
		}

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to save config:", err)
			return
		}

		cli.Success("Config updated:", key)
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	Run: func(cmd *cobra.Command, args []string) {
		var confirm string
		fmt.Print("This will reset all config. Continue? (y/N): ")
		fmt.Scanln(&confirm)

		if confirm != "y" && confirm != "Y" {
			cli.Warning("Aborted.")
			return
		}

		cfg := &config.Config{}

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to reset config:", err)
			return
		}

		cli.Success("Configuration reset to defaults.")
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configResetCmd)
}
