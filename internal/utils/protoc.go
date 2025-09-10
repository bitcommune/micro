package utils

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func CompileProtoFile(protoPath, goPackage string) error {
	// 获取proto文件所在目录
	protoDir := filepath.Dir(protoPath)

	// 构建protoc命令
	cmd := exec.Command("protoc",
		"--go_out", protoDir,
		"--go_opt", fmt.Sprintf("module=%s", goPackage),
		"--go-grpc_out", protoDir,
		"--go-grpc_opt", fmt.Sprintf("module=%s", goPackage),
		protoPath,
	)

	// 执行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("protoc execution failed: %v\nOutput: %s", err, output)
	}

	return nil
}
