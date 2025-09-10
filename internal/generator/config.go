package generator

import (
	"path/filepath"

	"github.com/bitcommune/micro/internal/utils"
)

func GenerateConfigManager(baseDir string) error {
	configData := map[string]interface{}{}

	// 生成配置加载器
	if err := utils.RenderTemplate("templates/config/config.go.tmpl", filepath.Join(baseDir, "pkg/config/config.go"), configData); err != nil {
		return err
	}

	// 生成配置结构
	if err := utils.RenderTemplate("templates/config/types.go.tmpl", filepath.Join(baseDir, "pkg/config/types.go"), configData); err != nil {
		return err
	}

	return nil
}
