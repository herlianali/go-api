package main

import "go-api/routes"

func main() {
	routes := routes.InitRoutes()
	routes.Run(":8080")
}
