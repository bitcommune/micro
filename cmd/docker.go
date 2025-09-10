package cmd

import (
	"fmt"
	"log"

	"github.com/bitcommune/micro/internal/generator"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Generate Docker configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.GenerateDockerConfig("."); err != nil {
			log.Fatalf("Error creating Docker configuration: %v", err)
		}

		fmt.Println("Docker configuration created successfully!")
	},
}
