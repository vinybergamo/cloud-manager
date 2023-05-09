package shell

import (
	terminal "github.com/nsf/termbox-go"
	"github.com/vinybergamo/cloud-deploy/database"
	sheelMenus "github.com/vinybergamo/cloud-deploy/shell/menus"
)

func Init() {
	options := []string{"Applications", "Databases", "Exit"}

	selected := sheelMenus.ShowMenuVertical(options)

	switch options[selected] {
	case "Applications":
		terminal.Close()
		sheelMenus.Applications()
		return
	case "Databases":
		terminal.Close()
		database.Create("teste")
		return
	case "Exit":
		terminal.Close()
		return
	}
}
