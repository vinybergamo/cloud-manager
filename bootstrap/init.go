package bootstrap

import (
	"flag"
	"fmt"

	"github.com/vinybergamo/cloud-manager/api"
	"github.com/vinybergamo/cloud-manager/shell"
	"github.com/vinybergamo/cloud-manager/vars"
)

func Init() {
	apiPtr := flag.Bool("api", false, "run api server")
	appPtr := flag.Bool("app", false, "run app")
	serverMode := flag.Bool("server-mode", false, "active server mode")

	flag.Parse()

	if *serverMode {
		vars.IsServerMode = true
	}

	if *apiPtr {
		api.Init()
	}

	if *appPtr {
		fmt.Println("Runing app")
	}

	if flag.NFlag() == 0 && !*serverMode {
		initShell()
	}

	if flag.NFlag() == 1 && *serverMode {
		initShell()
	}
}

func initShell() {
	shell.Init()
	vars.IsShellMode = true
}
