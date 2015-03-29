package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/firstrow/logvoyage/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := cli.NewApp()
	app.Name = "LogVoyage"
	app.Commands = []cli.Command{
		commands.StartBackendServer,
		commands.StartWebServer,
		commands.CreateUsersIndex,
		commands.DeleteIndex,
		commands.CreateIndex,
	}
	app.Run(os.Args)
}