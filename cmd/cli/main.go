package main

import (
	"github.com/deikioveca/betcrafter/internal/cli_app"
)

func main() {
	cli := cli_app.NewCliApp()
	cli.Run()
}