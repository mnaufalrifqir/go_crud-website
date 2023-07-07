package main

import "telkomsel/routes"

func main() {
	route := routes.StartRoute()
	route.Start(":8000")
}
