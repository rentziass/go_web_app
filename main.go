package main

import "gin/router"

func main() {
	r := router.RegisterRoutes()

	r.Run(":3000")
}
