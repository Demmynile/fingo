package api

import (
	"context"
	"net/http"

	db "github.com/demmynile/fingo/db/sqlc"
	"github.com/gin-gonic/gin"
)


type User struct {
	server *Server

}

func (u User) router(server *Server){
	u.server = server

	serverGroup := server.router.Group("/users")
	serverGroup.GET("" , u.listUsers)
	serverGroup.POST("" , u.createUser)
}

func (u *User) createUser(c *gin.Context) {

}

func (u *User) listUsers(c *gin.Context) {

	arg := db.ListUsersParams{
		Offset: 0,
		Limit: 10,
	}
	users , err := u.server.queries.ListUsers(context.Background() , arg)

	if err != nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
	}

	c.JSON(http.StatusOK, users)
}