package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestRenderTemplate(t *testing.T) {

	content, err := os.ReadFile("../../templates/project/main.go.tmpl")
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Println(string(content))

}
