package bootstrap

import (
	"flag"
	"fmt"

	"github.com/vinybergamo/cloud-deploy/shell"
)

func Init() {
	apiPtr := flag.Bool("api", false, "run api server")
	appPtr := flag.Bool("app", false, "run app")

	flag.Parse()

	if *apiPtr {
		fmt.Println("Runing api server")
	}

	if *appPtr {
		fmt.Println("Runing app")
	}

	if flag.NFlag() == 0 {
		shell.Init()
	}
}
