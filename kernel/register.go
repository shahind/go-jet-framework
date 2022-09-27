package kernel

import (
	"github.com/shahind/go-jet-framework/cli"
	"github.com/shahind/go-jet-framework/register"
)

var (
	// Commands will export all registered commands
	// The following map of interfaces expose all available method that can be used by Go-Jet CLI tool.
	// The map index determines the command that you've to run to for use the relative method.
	// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
	Commands = register.CommandRegister{
		"database:seed":      &cli.Seeder{},
		"show:commands":      &cli.ShowCommands{},
		"cmd:create":         &cli.CmdCreate{},
		"controller:create":  &cli.ControllerCreate{},
		"generate:key":       &cli.GenerateKey{},
		"middleware:create":  &cli.MiddlewareCreate{},
		"migration:create":   &cli.MigrationCreate{},
		"migration:rollback": &cli.MigrateRollback{},
		"migration:up":       &cli.MigrationUp{},
		"model:create":       &cli.ModelCreate{},
		"router:show":        &cli.RouterShow{},
		"service:create":     &cli.ServiceCreate{},
		"update":             &cli.UpdateAlfred{},
		// Here is where you've to register your custom controller
	}
	CommandServices = register.ServiceRegister{}
	Models          = register.ModelRegister{}
	Controllers     = register.ControllerRegister{}
	Middlewares     = register.MiddlewareRegister{}
	Router          []register.HTTPRouter
)
