package command

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type ControllerCreate struct {
	Signature   string
	Description string
}

func (c *ControllerCreate) Register() {
	c.Signature = "controller:create <name>"
	c.Description = "Create new controller"
}

func (c *ControllerCreate) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	cName := strings.Title(strings.ToLower(args))
	input, _ := ioutil.ReadFile(filepath.Join("../../raw/controller.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", gwf.GetDynamicPath("app/http/controller"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %sController has been created at %s", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
