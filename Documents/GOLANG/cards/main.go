package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var datas = []data{
	{Email: "echefu.amarachi@gmail.com", Password: "chicken123"},
	{Email: "apple.berry@gmail.com", Password: "nutty456"},
}

func getDatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)

}
func main() {
	router := gin.Default()
	router.GET("/datas", getDatas)
	router.Run("localhost: 8081")

}


