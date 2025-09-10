package test

import (
	"fmt"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	templatePath := "../templates/project/main.go.tmpl"
	//if err := utils.RenderTemplate("templates/project/main.go.tmpl", "", ""); err != nil {
	//	panic(err)
	//}

	// 读取模板文件
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(tmplContent))
}
