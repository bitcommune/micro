package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "micro",
	Short: "micro is a microservice code generator",
	Long: `A CLI tool for generating microservice project structure, 
RPC services, proto files, database integration, middleware, 
and Docker configuration.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Verbose mode enabled")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.micro.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// 添加子命令
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(protoCmd)
	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(databaseCmd)
}

func initConfig() {
	// 初始化配置
}
