package main

import (
	"github.com/agorago/storeservice"
	"github.com/agorago/wego/cmd"
)

func main() {
	cmd.SwaggerMain(storeservice.Initializers...)
}
