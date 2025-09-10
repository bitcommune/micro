package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"unicode"
)

func RenderTemplate(templatePath, outputPath string, data interface{}) error {
	exPath, err := os.Getwd()
	if err != nil {
		return err
	}
	// 读取模板文件
	tmplContent, err := os.ReadFile(fmt.Sprintf("%s/%v", filepath.Dir(exPath), templatePath))
	if err != nil {
		return err
	}

	// 解析模板
	tmpl, err := template.New("template").Parse(string(tmplContent))
	if err != nil {
		return err
	}

	// 执行模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// 写入文件
	return CreateFile(outputPath, buf.String())
}

func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func ToCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	result := make([]rune, 0, len(s))
	nextUpper := true

	for _, r := range s {
		if r == '_' {
			nextUpper = true
			continue
		}

		if nextUpper {
			result = append(result, unicode.ToUpper(r))
			nextUpper = false
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}

	return string(result)
}
