package sheelMenus

import (
	terminal "github.com/nsf/termbox-go"
	"github.com/vinybergamo/cloud-deploy/apps"
	"github.com/vinybergamo/cloud-deploy/domains"
	"github.com/vinybergamo/cloud-deploy/letsencrypt"
	"github.com/vinybergamo/cloud-deploy/proxy"
	"github.com/vinybergamo/cloud-deploy/utils"
)

func Applications() {
	options := []string{"Create", "Deploy", "Destroy", "Add Domain", "List", "Check", "Exit"}

	selected := ShowMenuVertical(options)

	switch options[selected] {
	case "Create":
		terminal.Close()
		app := utils.ReadInput("Enter application name")
		apps.Create(app)
		return
	case "Deploy":
		terminal.Close()
		app := utils.ReadInput("Enter application name")
		git := utils.ReadInput("Enter Git repository")
		port := utils.ReadInput("Enter application port")
		apps.Deploy(app, git, port)
		return
	case "Add Domain":
		terminal.Close()
		app := utils.ReadInput("Enter application name")
		domain := utils.ReadInput("Enter domain name")
		domains.Add(app, domain)
		letsencrypt.Enable(app)
		return
	case "Destroy":
		app := utils.ReadInput("Enter application name")
		proxy.ClearConfig(app)
		apps.Destroy(app)
		return
	case "List":
		terminal.Close()
		apps.List()
		return
	case "Check":
		terminal.Close()
		app := utils.ReadInput("Enter application name")
		apps.ChecksRun(app)
		return
	case "Exit":
		terminal.Close()
		return
	}
}
