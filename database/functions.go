package database

import (
	"github.com/vinybergamo/cloud-deploy/utils"
)

func Create(n string) string {
	utils.Logger("yellow", "Creating database", n+"...")
	_, err := utils.ExecCommand("postgres:create", n)
	if err != nil {
		utils.LoggerError(err.Error())
	}

	utils.Logger("green", "Database", n, "created")
	return "create"
}
