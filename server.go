package main

import (
	"gatewayCore/http"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {

	server := gin.Default()

	// ===== Train GraphQL Module ===== //

	server.GET("/train/playground", http.PlaygroundHandler("/train"))

	server.POST("/train", http.GraphQLHandler("/train"))

	// ================================ //

	// ===== Flight GraphQL Module ==== //

	server.GET("/flight/playground", http.PlaygroundHandler("/flight"))

	server.POST("/flight", http.GraphQLHandler("/flight"))

	// ================================= //

	server.Run(defaultPort)

}