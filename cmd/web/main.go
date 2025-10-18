package main

import (
	"github.com/deikioveca/betcrafter/internal/web_app"
)

func main() {
	webApp := web_app.NewWebApp()
	webApp.Run(":8080")
}