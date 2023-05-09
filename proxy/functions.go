package proxy

import (
	"github.com/vinybergamo/cloud-manager/utils"
)

func Set(n, p string) {
	utils.Logger("yellow", "Setting proxy for", n, "to", p, "...")
	_, err := utils.ExecCommand("proxy:ports-set", n, "http:80:"+p)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "Proxy set successfully")
}

func ClearConfig(n string) {
	utils.Logger("yellow", "Clearing proxy config for", n, "...")
	_, err := utils.ExecCommand("proxy:clear-config", n)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "Proxy config cleared successfully")
}
