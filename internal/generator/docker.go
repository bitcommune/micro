package generator

import (
	"path/filepath"

	"github.com/bitcommune/micro/internal/utils"
)

func GenerateDockerConfig(baseDir string) error {
	// 生成Dockerfile
	dockerfileData := map[string]interface{}{
		"ProjectName": filepath.Base(baseDir),
	}

	if err := utils.RenderTemplate("templates/project/dockerfile.tmpl", filepath.Join(baseDir, "Dockerfile"), dockerfileData); err != nil {
		return err
	}

	// 生成docker-compose.yml
	dockerComposeData := map[string]interface{}{
		"ProjectName": filepath.Base(baseDir),
	}

	if err := utils.RenderTemplate("templates/project/docker-compose.yaml.tmpl", filepath.Join(baseDir, "docker-compose.yaml"), dockerComposeData); err != nil {
		return err
	}

	return nil
}
