package main

import router "example/mongo-go/router"

func main() {
	r := router.Routers()
	r.Run("localhost:8080")
}
