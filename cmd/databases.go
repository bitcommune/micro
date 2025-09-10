package cmd

import (
	"fmt"
	"log"

	"github.com/bitcommune/micro/internal/generator"
	"github.com/bitcommune/micro/internal/types"
	"github.com/spf13/cobra"
)

//var dbType string

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Generate database integration",
	Run: func(cmd *cobra.Command, args []string) {
		config := types.DatabaseConfig{
			Type: dbType,
			Models: []types.Model{
				{
					Name: "User",
					Fields: []types.Field{
						{Name: "ID", Type: "uint"},
						{Name: "Name", Type: "string"},
						{Name: "Email", Type: "string"},
					},
				},
			},
		}

		if err := generator.GenerateDatabaseIntegration(config, "."); err != nil {
			log.Fatalf("Error creating database integration: %v", err)
		}

		fmt.Printf("Database integration for %s created successfully!\n", dbType)
	},
}

func init() {
	databaseCmd.Flags().StringVar(&dbType, "type", "postgres", "Database type (postgres, mysql, sqlite)")
}
