package main

import (
	"fmt"

	"github.com/Bhargav16exd/serverctl/commands"
	"github.com/Bhargav16exd/serverctl/generators"
)

func main() {

	commands.CheckNginxInstallation()

	path, err := commands.FetchNginxConfPath()

	if err != nil {
		fmt.Println(err)
	}

	sitesAvailableDir, _ := commands.CheckCreateSitesDir(path)

	generators.GenerateApiConfig("somename", sitesAvailableDir, "3000", "localhost")
}
