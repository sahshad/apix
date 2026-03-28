package cmd

import (
	"github.com/sahshad/apix/internal/cli"
	"github.com/sahshad/apix/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage API environments",
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available environments",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		if len(cfg.Environments) == 0 {
			cli.Warning("No environments configured.")
			return
		}

		fmt.Println("Available environments:")
		for name, url := range cfg.Environments {
			if name == cfg.CurrentEnv {
				fmt.Println("→", name, "(", url, ")")
			} else {
				fmt.Println("-", name, "(", url, ")")
			}
		}
	},
}

var envUseCmd = &cobra.Command{
	Use:   "use [environment]",
	Short: "Switch to a specific environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		url, ok := cfg.Environments[envName]
		if !ok {
			cli.Error("Environment not found:", envName)
			return
		}

		cfg.CurrentEnv = envName
		cfg.BaseURL = url

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to save config:", err)
			return
		}

		cli.Success("Switched to environment:", envName)
	},
}

var envSetCmd = &cobra.Command{
	Use:   "set NAME=URL",
	Short: "Add or update an environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		if cfg.Environments == nil {
			cfg.Environments = make(map[string]string)
		}

		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			cli.Error("Invalid format. Use NAME=URL")
			return
		}

		name := strings.TrimSpace(parts[0])
		url := strings.TrimSpace(parts[1])

		if name == "" || url == "" {
			cli.Error("Name and URL cannot be empty")
			return
		}

		cfg.Environments[name] = url

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to save config:", err)
			return
		}

		cli.Success("Environment saved:", name)
	},
}

var envDeleteCmd = &cobra.Command{
	Use:   "delete [environment]",
	Short: "Delete an environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			cli.Error("Failed to load config:", err)
			return
		}

		if _, ok := cfg.Environments[envName]; !ok {
			cli.Warning("Environment not found:", envName)
			return
		}

		delete(cfg.Environments, envName)

		if cfg.CurrentEnv == envName {
			cfg.CurrentEnv = ""
			cfg.BaseURL = ""
		}

		if err := config.SaveConfig(cfg); err != nil {
			cli.Error("Failed to save config:", err)
			return
		}

		cli.Success("Environment deleted:", envName)
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envUseCmd)
	envCmd.AddCommand(envSetCmd)
	envCmd.AddCommand(envDeleteCmd)
}
