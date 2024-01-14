package main

import (
	"contractTest/pkg/local/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users/:userId", server.GetUserByID)
	router.Run(":9001")

	// Uncomment this code if you want to run client code

	// user, err := client.GetUserByID("localhost:9001", "1")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

}
