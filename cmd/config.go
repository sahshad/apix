package cmd

import (
	"apix/internal/config"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current CLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		b, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Println(string(b))
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [KEY=VALUE]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		var key, value string
		fmt.Sscanf(args[0], "%[^=]=%s", &key, &value)

		switch key {
		case "default_output":
			// example config option
			fmt.Println("Set", key, "=", value)
		default:
			fmt.Println("Unknown config key:", key)
			return
		}

		config.SaveConfig(cfg)
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &config.Config{}
		config.SaveConfig(cfg)
		fmt.Println("Configuration reset to defaults.")
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configResetCmd)
}