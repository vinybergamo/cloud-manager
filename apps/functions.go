package apps

import (
	"fmt"
	"strings"

	"github.com/vinybergamo/cloud-manager/git"
	"github.com/vinybergamo/cloud-manager/letsencrypt"
	"github.com/vinybergamo/cloud-manager/proxy"
	"github.com/vinybergamo/cloud-manager/utils"
	"github.com/vinybergamo/cloud-manager/vars"
)

func Exists(n string) bool {
	utils.Logger("yellow", "Checking app...")
	_, err := utils.ExecCommand("apps:exists", n)
	return err == nil
}

func Create(n string) (string, error) {
	exists := Exists(n)
	if exists {
		utils.LoggerError("App already exists")

		r := appsResponse{
			Error:   true,
			Message: "App already exists",
			Status:  "Error",
			Service: "Apps",
			Name:    n,
		}
		return utils.CreateJSONResponse(r), fmt.Errorf("app already exists")
	}

	utils.Logger("yellow", "Creating app... ")

	_, err := utils.ExecCommand("apps:create", n)
	if err != nil {
		utils.LoggerError(err.Error())

		r := appsResponse{
			Error:   true,
			Message: "Error creating app: " + err.Error(),
			Status:  "Error",
			Service: "Apps",
			Name:    n,
		}
		return utils.CreateJSONResponse(r), err
	}

	utils.Logger("green", n, "created")
	r := appsResponse{
		Message: "App created",
		Status:  "Created",
		Service: "Apps",
		Name:    n,
	}

	return utils.CreateJSONResponse(r), nil

}

func Destroy(n string) (string, error) {
	exists := Exists(n)
	if !exists {
		utils.LoggerError("App does not exists")

		r := appsResponse{
			Error:   true,
			Message: "App does not exists",
			Status:  "Error",
			Service: "Apps",
			Name:    n,
		}

		return utils.CreateJSONResponse(r), fmt.Errorf("app does not exists")
	}

	utils.Logger("yellow", "Destroying app... ")

	_, err := utils.ExecCommand("apps:destroy", n, "--force")
	if err != nil {
		utils.LoggerError(err.Error())

		r := appsResponse{
			Error:   true,
			Message: "Error destroying app: " + err.Error(),
			Status:  "Error",
			Service: "Apps",
			Name:    n,
		}

		return utils.CreateJSONResponse(r), err
	}

	utils.Logger("green", "App", n, "destroyed")

	r := appsResponse{
		Error:   false,
		Message: "App destroyed",
		Status:  "Destroyed",
		Service: "Apps",
		Name:    n,
	}

	return utils.CreateJSONResponse(r), nil
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

func Deploy(n, u, p string) (string, error) {
	_, err := Create(n)
	if err != nil {
		r := appsResponse{
			Status:  "Error",
			Message: err.Error(),
			Service: "Apps",
			Name:    n,
			Error:   true,
		}

		return utils.CreateJSONResponse(r), err
	}

	git.Sync(n, u)
	Build(n)
	proxy.Set(n, p)
	letsencrypt.Enable(n)
	letsencrypt.EnableCronjob()
	utils.Logger("green", "App", n, "deployed")
	_, err = utils.CheckHTTPSDomain(n + "." + vars.DefaultDomain)
	utils.Logger("green", "Application available at:")
	if err != nil {
		utils.Logger("green", "http://"+n+".")
		r := appsResponse{
			Status:  "Deployed",
			Message: "Application deployed",
			Service: "Apps",
			Name:    n,
			Url:     "http://" + n + ".",
			Error:   false,
		}
		return utils.CreateJSONResponse(r), nil
	}
	utils.Logger("green", "https://"+n+"."+vars.DefaultDomain)
	r := appsResponse{
		Status:  "Deployed",
		Message: "Application deployed",
		Service: "Apps",
		Name:    n,
		Url:     "https://" + n + "." + vars.DefaultDomain,
		Error:   false,
	}
	return utils.CreateJSONResponse(r), nil
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
