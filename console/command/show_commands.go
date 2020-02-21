package command

import (
	"os"
	"reflect"

	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/olekukonko/tablewriter"
)

// ShowCommands will show all registered commands
type ShowCommands struct {
	Signature   string
	Description string
}

// Register this command
func (c *ShowCommands) Register() {
	c.Signature = "show:commands"
	c.Description = "Show Go-Web commands list"
}

// Run this command
func (c *ShowCommands) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	var data [][]string

	for _, c := range console {
		m := reflect.ValueOf(c).MethodByName("Register")
		m.Call([]reflect.Value{})

		cmd := reflect.ValueOf(c).Elem()

		signature := cmd.FieldByName("Signature").String()
		description := cmd.FieldByName("Description").String()

		data = append(data, []string{signature, description})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMMAND", "DESCRIPTION"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
