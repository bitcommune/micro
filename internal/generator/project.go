package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bitcommune/micro/internal/types"
	"github.com/bitcommune/micro/internal/utils"
)

func GenerateProject(config types.ProjectConfig) error {
	baseDir := config.Name
	dirs := []string{
		"api/proto",
		"cmd/server",
		"cmd/client",
		"internal/service",
		"internal/handler",
		"internal/repository",
		"internal/middleware",
		"pkg/pb",
		"configs",
		"scripts",
		"migrations",
		"deployments",
	}

	// 创建目录
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(baseDir, dir), 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// 创建go.mod文件
	if err := createGoMod(baseDir, config); err != nil {
		return err
	}

	// 创建main.go文件
	if err := createMainFile(baseDir, config); err != nil {
		return err
	}

	// 创建配置文件
	if err := createConfigFile(baseDir, config); err != nil {
		return err
	}

	// 创建Makefile
	if err := createMakefile(baseDir, config); err != nil {
		return err
	}

	// 如果启用了数据库，生成数据库集成
	if config.WithDB {
		dbConfig := types.DatabaseConfig{
			Type: config.DBType,
			Models: []types.Model{
				{
					Name: "Example",
					Fields: []types.Field{
						{Name: "ID", Type: "uint"},
						{Name: "Name", Type: "string"},
					},
				},
			},
		}
		if err := GenerateDatabaseIntegration(dbConfig, baseDir); err != nil {
			return err
		}
	}

	// 如果启用了Docker，生成Docker配置
	if config.WithDocker {
		if err := GenerateDockerConfig(baseDir); err != nil {
			return err
		}
	}

	// 如果启用了中间件，生成中间件
	if config.WithMiddleware {
		if err := GenerateMiddleware(baseDir); err != nil {
			return err
		}
	}

	// 如果启用了测试，生成测试模板
	if config.WithTests {
		if err := GenerateTestTemplates(baseDir); err != nil {
			return err
		}
	}

	return nil
}

func createGoMod(baseDir string, config types.ProjectConfig) error {
	goModContent := fmt.Sprintf(`module %s

go 1.21

require (
	github.com/spf13/viper v1.16.0
	google.golang.org/grpc v1.58.0
	google.golang.org/protobuf v1.31.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.4
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-tobuilder v1.9.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
)
`, config.GoModule)

	return utils.CreateFile(filepath.Join(baseDir, "go.mod"), goModContent)
}

func createMainFile(baseDir string, config types.ProjectConfig) error {
	templateData := map[string]interface{}{
		"GoModule":       config.GoModule,
		"WithDB":         config.WithDB,
		"WithMiddleware": config.WithMiddleware,
	}

	return utils.RenderTemplate("templates/project/main.go.tmpl", filepath.Join(baseDir, "cmd/server/main.go"), templateData)
}

func createConfigFile(baseDir string, config types.ProjectConfig) error {
	templateData := map[string]interface{}{
		"WithDB": config.WithDB,
		"DBType": config.DBType,
	}

	return utils.RenderTemplate("templates/project/config.yaml.tmpl", filepath.Join(baseDir, "configs/config.yaml"), templateData)
}

func createMakefile(baseDir string, config types.ProjectConfig) error {
	templateData := map[string]interface{}{
		"ProjectName": config.Name,
		"WithDB":      config.WithDB,
		"WithDocker":  config.WithDocker,
	}

	return utils.RenderTemplate("templates/project/Makefile.tmpl", filepath.Join(baseDir, "Makefile"), templateData)
}
