package cmd

import (
	"fmt"
	"log"

	"github.com/bitcommune/micro/internal/generator"
	"github.com/bitcommune/micro/internal/types"
	"github.com/spf13/cobra"
)

var (
	withDB         bool
	dbType         string
	withDocker     bool
	withTests      bool
	withMiddleware bool
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new microservice project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		config := types.ProjectConfig{
			Name:           projectName,
			Version:        "1.0.0",
			GoModule:       fmt.Sprintf("github.com/your-username/%s", projectName),
			WithDB:         withDB,
			DBType:         dbType,
			WithDocker:     withDocker,
			WithTests:      withTests,
			WithMiddleware: withMiddleware,
		}

		if err := generator.GenerateProject(config); err != nil {
			log.Fatalf("Error creating project: %v", err)
		}

		fmt.Printf("Project %s created successfully!\n", projectName)
	},
}

func init() {
	initCmd.Flags().BoolVar(&withDB, "with-db", false, "Include database integration")
	initCmd.Flags().StringVar(&dbType, "db-type", "postgres", "Database type (postgres, mysql, sqlite)")
	initCmd.Flags().BoolVar(&withDocker, "with-docker", false, "Include Docker configuration")
	initCmd.Flags().BoolVar(&withTests, "with-tests", false, "Include test templates")
	initCmd.Flags().BoolVar(&withMiddleware, "with-middleware", false, "Include middleware")
}
