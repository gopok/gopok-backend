package main

import (
	_ "github.com/alufers/fw/pkg/auth"
	"github.com/alufers/fw/pkg/core"
)

func main() {
	app := core.Application{}
	app.Run()
}
