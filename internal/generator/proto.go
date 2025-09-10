package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bitcommune/micro/internal/types"
	"github.com/bitcommune/micro/internal/utils"
)

func GenerateProtoFile(config types.ProtoConfig, outputPath string) error {
	templateData := map[string]interface{}{
		"PackageName": config.PackageName,
		"GoPackage":   config.GoPackage,
		"ServiceName": config.ServiceName,
		"Messages":    config.Messages,
		"RPCs":        config.RPCs,
	}

	// 确保输出目录存在
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return utils.RenderTemplate("templates/proto/service.proto.tmpl", outputPath, templateData)
}

func CompileProto(protoPath, goPackage string) error {
	return utils.CompileProtoFile(protoPath, goPackage)
}
