package git

import (
	"github.com/vinybergamo/cloud-deploy/utils"
)

func Sync(n, u string) {
	utils.Logger("yellow", "Syncing", n, "with", u+"...")
	_, err := utils.ExecCommand("git:sync", n, u)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "Synced", n, "with", u)
}
