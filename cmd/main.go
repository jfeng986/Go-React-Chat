package main

import (
	"Go-React-Chat/router"
)

func main() {
	router.InitRouter()
	router.Start(":8080")
}
