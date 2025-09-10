package cmd

import (
	"fmt"
	"log"

	"github.com/bitcommune/micro/internal/generator"
	"github.com/bitcommune/micro/internal/types"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service [service-name]",
	Short: "Create a new RPC service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]

		config := types.ServiceConfig{
			Name:        serviceName,
			PackageName: fmt.Sprintf("github.com/bitcommune/%v/pkg/pb", serviceName),
			RPCs: []types.RPC{
				{Name: "Example", Request: "ExampleRequest", Response: "ExampleResponse"},
			},
		}

		if err := generator.GenerateService(config, "."); err != nil {
			log.Fatalf("Error creating service: %v", err)
		}

		fmt.Printf("Service %s created successfully!\n", serviceName)
	},
}
