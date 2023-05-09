package apps

import (
	"fmt"
	"os"
	"strings"

	"github.com/vinybergamo/cloud-deploy/git"
	"github.com/vinybergamo/cloud-deploy/letsencrypt"
	"github.com/vinybergamo/cloud-deploy/proxy"
	"github.com/vinybergamo/cloud-deploy/utils"
	"github.com/vinybergamo/cloud-deploy/vars"
)

func Exists(n string) bool {
	utils.Logger("yellow", "Checking app...")
	_, err := utils.ExecCommand("apps:exists", n)
	return err == nil
}

func Create(n string) {
	exists := Exists(n)
	if exists {
		utils.LoggerError("App already exists")
	}

	utils.Logger("yellow", "Creating app... ")

	_, err := utils.ExecCommand("apps:create", n)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	utils.Logger("green", n, "created")
}

func Destroy(n string) {
	exists := Exists(n)
	if !exists {
		utils.LoggerError("App does not exists")
	}

	utils.Logger("yellow", "Destroying app... ")

	_, err := utils.ExecCommand("apps:destroy", n, "--force")
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "App", n, "destroyed")
}

func Build(n string) {
	exists := Exists(n)
	if !exists {
		utils.LoggerError("App does not exists")
	}

	utils.Logger("yellow", "Building app... ")

	_, err := utils.ExecCommand("ps:rebuild", n)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "App", n, "built")
}

func Update(n string) {
	exists := Exists(n)
	if !exists {
		utils.LoggerError("App does not exists")
	}

	utils.Logger("yellow", "Updating app...")

	Build(n)

	utils.Logger("green", "App", n, "updated")
}

func List() {
	apps, err := utils.ExecCommand("apps:list")
	if err != nil {
		utils.LoggerError(err.Error())
	}

	apps = strings.Replace(apps, "=====> My Apps\n", "", 1)
	apps = strings.TrimSpace(apps)

	arrayApps := strings.Split(apps, "\n")
	for i, app := range arrayApps {
		arrayApps[i] = strings.TrimSpace(app)
	}

	utils.Logger("green", apps)
}

func Deploy(n, u, p string) {
	Create(n)
	git.Sync(n, u)
	Build(n)
	proxy.Set(n, p)
	letsencrypt.Enable(n)
	letsencrypt.EnableCronjob()
	utils.Logger("green", "App", n, "deployed")
	_, err := utils.CheckHTTPSDomain(n + "." + vars.DefaultDomain)
	utils.Logger("green", "Application available at:")
	if err != nil {
		utils.Logger("green", "http://"+n+".")
		os.Exit(0)
	}
	utils.Logger("green", "https://"+n+"."+vars.DefaultDomain)
	os.Exit(0)
}

func ChecksRun(n string) string {
	utils.Logger("yellow", "Checking app...")
	output, err := utils.ExecCommand("checks:run", n)
	if err != nil {
		utils.LoggerError(err.Error())
		return err.Error()
	}

	utils.Logger("green", output)
	return output
}
