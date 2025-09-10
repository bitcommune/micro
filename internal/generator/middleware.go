package generator

import (
	"path/filepath"

	"github.com/bitcommune/micro/internal/utils"
)

func GenerateMiddleware(baseDir string) error {
	// 生成日志中间件
	if err := utils.RenderTemplate("templates/middleware/logging.go.tmpl", filepath.Join(baseDir, "internal/middleware/logging.go"), nil); err != nil {
		return err
	}

	// 生成认证中间件
	if err := utils.RenderTemplate("templates/middleware/auth.go.tmpl", filepath.Join(baseDir, "internal/middleware/auth.go"), nil); err != nil {
		return err
	}

	// 生成恢复中间件
	if err := utils.RenderTemplate("templates/middleware/recovery.go.tmpl", filepath.Join(baseDir, "internal/middleware/recovery.go"), nil); err != nil {
		return err
	}

	return nil
}
