package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id" pact:"example=1"`
	FirstName string `json:"firstName" pact:"example=Alice"`
	LastName  string `json:"lastName" pact:"example=Doe"`
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userId")

	ctx.JSON(http.StatusOK, User{
		ID:        id,
		FirstName: fmt.Sprintf("first%s", id),
		LastName:  fmt.Sprintf("last%s", id),
	})
}
