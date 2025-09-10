package generator

import (
	"path/filepath"

	"github.com/bitcommune/micro/internal/utils"
)

func GenerateTestTemplates(baseDir string) error {
	// 生成服务测试模板
	serviceTestData := map[string]interface{}{}
	serviceTestPath := filepath.Join(baseDir, "internal/service/service_test.go")
	if err := utils.RenderTemplate("templates/test/service_test.go.tmpl", serviceTestPath, serviceTestData); err != nil {
		return err
	}

	// 生成处理器测试模板
	handlerTestData := map[string]interface{}{}
	handlerTestPath := filepath.Join(baseDir, "internal/handler/handler_test.go")
	if err := utils.RenderTemplate("templates/test/handler_test.go.tmpl", handlerTestPath, handlerTestData); err != nil {
		return err
	}

	// 生成存储库测试模板
	repoTestData := map[string]interface{}{}
	repoTestPath := filepath.Join(baseDir, "internal/repository/repository_test.go")
	if err := utils.RenderTemplate("templates/test/repository_test.go.tmpl", repoTestPath, repoTestData); err != nil {
		return err
	}

	return nil
}
