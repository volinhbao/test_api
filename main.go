package main

import (

	"log"
	"github.com/gin-gonic/gin"

	routes "github.com/volinhbao/test_api/routes"
)

func main() {
	
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}