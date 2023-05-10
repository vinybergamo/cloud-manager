package vars

import "os"

var Server = os.Getenv("SERVER")
var DefaultDomain = "vinybergamo.cloud"

var IsShellMode bool = false
