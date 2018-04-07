package main

import (
	_ "github.com/gopok/gopok-backend/pkg/auth"
	_ "github.com/gopok/gopok-backend/pkg/blog"
	"github.com/gopok/gopok-backend/pkg/core"
)

func main() {
	app := core.Application{}
	app.Run()
}
