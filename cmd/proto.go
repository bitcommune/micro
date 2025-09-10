package cmd

import (
	"fmt"
	"github.com/bitcommune/micro/internal/utils"
	"log"

	"github.com/bitcommune/micro/internal/generator"
	"github.com/bitcommune/micro/internal/types"
	"github.com/spf13/cobra"
)

var (
	protoPackage string
	goPackage    string
	compileProto bool
)

var protoCmd = &cobra.Command{
	Use:   "proto [proto-name]",
	Short: "Create a new proto file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		protoName := args[0]

		config := types.ProtoConfig{
			PackageName: protoPackage,
			GoPackage:   goPackage,
			ServiceName: utils.ToCamelCase(protoName),
			Messages: []types.Message{
				{
					Name: "ExampleRequest",
					Fields: []types.Field{
						{Type: "string", Name: "name", Tag: 1},
					},
				},
				{
					Name: "ExampleResponse",
					Fields: []types.Field{
						{Type: "string", Name: "message", Tag: 1},
					},
				},
			},
			RPCs: []types.RPC{
				{Name: "Example", Request: "ExampleRequest", Response: "ExampleResponse"},
			},
		}

		outputPath := fmt.Sprintf("api/proto/%s.proto", protoName)

		if err := generator.GenerateProtoFile(config, outputPath); err != nil {
			log.Fatalf("Error creating proto file: %v", err)
		}

		if compileProto {
			if err := generator.CompileProto(outputPath, goPackage); err != nil {
				log.Fatalf("Error compiling proto file: %v", err)
			}
			fmt.Println("Proto file compiled successfully!")
		}

		fmt.Printf("Proto file %s created successfully!\n", protoName)
	},
}

func init() {
	protoCmd.Flags().StringVarP(&protoPackage, "package", "p", "myapp", "Package name for the proto file")
	protoCmd.Flags().StringVarP(&goPackage, "gopackage", "g", "github.com/your-username/myapp/pkg/pb", "Go package path")
	protoCmd.Flags().BoolVarP(&compileProto, "compile", "c", false, "Compile proto file after generation")
}
