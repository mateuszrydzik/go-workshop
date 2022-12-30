package json_parse

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type input struct{
	name string `json:"name"`
	favNumber int `json:"favNumber"`
}

func post(c *gin.Context) {
	var json input

	if err := context.BindJSON(&json); err != nil {
		return
	}
	c.IntendedJSON(http.StatusCreated)
}

func main()  {
	router := gin.Default()
	router.Run("localhost:8080")
	router.POST("/data", post)
}

func homePage(response http.ResponseWriter, r *http.request)