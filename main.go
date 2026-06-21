package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Bhargav16exd/nginxctl/ui"
)

func main() {
	if os.Geteuid() == 0 {
		p := tea.NewProgram(ui.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("THE UTILITY REQUIRES SUDO PERMISSIONS, PLEASE GIVE SUDO PERMISSIONS")
		os.Exit(1)
	}
}
