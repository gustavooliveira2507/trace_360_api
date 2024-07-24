package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var bills = []bill{
	{
		ID:       "1",
		BillType: bill_type{Id: "1", Description: "Teste"},
	},
}

func getBills() []bill {
	return bills
}
func main() {
	router := gin.Default()
	router.GET(
		"/bill",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, getBills())
		},
	)

	router.Run("localhost:8080")
}
