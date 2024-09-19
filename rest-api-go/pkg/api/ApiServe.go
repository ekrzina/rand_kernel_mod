package apiserve

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	Port string
}

/* To be replaced with kernel module */
func getRandomNumber(c *gin.Context) {
	randomNumber := rand.Intn(10000)
	c.JSON(http.StatusOK, gin.H{
		"randomNumber": randomNumber,
	})
}

/* Serves REST API on specified port */
func (t *ApiHandler) HandleRequests() {
	router := gin.Default()
	router.GET("/randnumber", getRandomNumber)

	fmt.Printf("Server started on port %s...\n", t.Port)
	err := router.Run(":" + t.Port)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
