package routes

import (

)
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": Welcome to API,
	})
	return
}