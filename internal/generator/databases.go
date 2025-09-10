package generator

import (
	"fmt"
	"path/filepath"

	"github.com/bitcommune/micro/internal/types"
	"github.com/bitcommune/micro/internal/utils"
)

func GenerateDatabaseIntegration(config types.DatabaseConfig, baseDir string) error {
	// 生成数据库连接配置
	dbConfigData := map[string]interface{}{
		"DBType": config.Type,
	}

	if err := utils.RenderTemplate("templates/database/config.go.tmpl", filepath.Join(baseDir, "internal/repository/database.go"), dbConfigData); err != nil {
		return err
	}

	// 生成模型
	for _, model := range config.Models {
		modelData := map[string]interface{}{
			"Model": model,
		}

		modelPath := filepath.Join(baseDir, "internal/repository/models", fmt.Sprintf("%s.go", utils.ToSnakeCase(model.Name)))
		if err := utils.RenderTemplate("templates/database/model.go.tmpl", modelPath, modelData); err != nil {
			return err
		}

		// 生成存储库接口和实现
		repoData := map[string]interface{}{
			"Model": model,
		}

		repoPath := filepath.Join(baseDir, "internal/repository", fmt.Sprintf("%s_repository.go", utils.ToSnakeCase(model.Name)))
		if err := utils.RenderTemplate("templates/database/repository.go.tmpl", repoPath, repoData); err != nil {
			return err
		}
	}

	// 生成迁移文件
	migrationData := map[string]interface{}{
		"Models": config.Models,
		"DBType": config.Type,
	}

	migrationPath := filepath.Join(baseDir, "migrations", "001_init.sql")
	if err := utils.RenderTemplate("templates/database/migration.sql.tmpl", migrationPath, migrationData); err != nil {
		return err
	}

	return nil
}
