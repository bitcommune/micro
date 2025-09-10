package generator

import (
	"fmt"
	"path/filepath"

	"github.com/bitcommune/micro/internal/types"
	"github.com/bitcommune/micro/internal/utils"
)

func GenerateService(config types.ServiceConfig, baseDir string) error {
	// 生成服务接口和实现
	serviceTemplateData := map[string]interface{}{
		"ServiceName": config.ServiceName,
		"PackageName": config.PackageName,
		"RPCs":        config.RPCs,
	}

	servicePath := filepath.Join(baseDir, "internal/service", fmt.Sprintf("%s.go", utils.ToSnakeCase(config.ServiceName)))
	if err := utils.RenderTemplate("templates/service/service.go.tmpl", servicePath, serviceTemplateData); err != nil {
		return fmt.Errorf("failed to generate service: %v", err)
	}

	// 生成处理器
	handlerTemplateData := map[string]interface{}{
		"ServiceName": config.ServiceName,
		"PackageName": config.PackageName,
		"RPCs":        config.RPCs,
	}

	handlerPath := filepath.Join(baseDir, "internal/handler", fmt.Sprintf("%s_handler.go", utils.ToSnakeCase(config.ServiceName)))
	if err := utils.RenderTemplate("templates/service/handler.go.tmpl", handlerPath, handlerTemplateData); err != nil {
		return fmt.Errorf("failed to generate handler: %v", err)
	}

	return nil
}
