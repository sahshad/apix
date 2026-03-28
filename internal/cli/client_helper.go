package cli

import (
	"apix/internal/client"
	"apix/internal/config"
	"fmt"
	"os"
)

func GetClient() *client.APIClient {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	clnt, err := client.NewClient(cfg)
	if err != nil {
		fmt.Println("Failed to init client:", err)
		os.Exit(1)
	}

	return clnt
}