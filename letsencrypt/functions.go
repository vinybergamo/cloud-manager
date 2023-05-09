package letsencrypt

import "github.com/vinybergamo/cloud-deploy/utils"

func Enable(n string) {
	utils.Logger("yellow", "Enabling Let's Encrypt")

	_, err := utils.ExecCommand("letsencrypt:enable", n)
	if err != nil {
		utils.Logger("red", "Error enabling Let's Encrypt: ", err.Error())
		return
	}

	utils.Logger("green", "Let's Encrypt enabled")
}

func EnableCronjob() {
	utils.Logger("yellow", "Enabling Let's Encrypt cronjob...")
	_, err := utils.ExecCommand("letsencrypt:cron-job", "--add")
	if err != nil {
		utils.Logger("red", "Error enabling Let's Encrypt cronjob")
		return
	}

	utils.Logger("green", "Let's Encrypt cronjob enabled")
}
