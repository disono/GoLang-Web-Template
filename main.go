package main

import (
	"github.com/disono/GoLang-Web-Template/app/presenters/routes"
	"github.com/disono/GoLang-Web-Template/app/framework/validator"
)

func init()  {
	validator.Run()
}

func main() {
	routes.InitializeRoutes()
}
