package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestRenderTemplate(t *testing.T) {

	exPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("可执行文件所在目录 (os.Executable()):", filepath.Dir(exPath))
	content, err := os.ReadFile(fmt.Sprintf("%s/templates/project/main.go.tmpl", filepath.Dir(exPath)))
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Println(string(content))

}
