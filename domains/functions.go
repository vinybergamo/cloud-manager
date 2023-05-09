package domains

import (
	"github.com/vinybergamo/cloud-manager/apps"
	"github.com/vinybergamo/cloud-manager/utils"
	"github.com/vinybergamo/cloud-manager/vars"
)

func Add(a string, d string) {
	dIp, err := utils.CheckHTTPDomain(d)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	bIp, err := utils.CheckHTTPDomain("cname." + vars.DefaultDomain)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	if dIp != bIp {
		utils.LoggerError("Domain IP does not match the domain name")
	}

	exists := apps.Exists(a)
	if !exists {
		utils.LoggerError("App not found")
	}

	utils.Logger("yellow", "Adding domains...")
	_, err = utils.ExecCommand("domains:add", a, d)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "Domains", d, "is now available for", a)
}
