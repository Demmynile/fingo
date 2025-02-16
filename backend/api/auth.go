package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/demmynile/fingo/utils"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	server *Server
}

func (a Auth) router (server *Server) {
	a.server = server

	serverGroup := server.router.Group("/auth")
	serverGroup.POST("login" ,  a.login)
}

func (a Auth) login(c *gin.Context){
	user := new (UserParams)
	if err :=  c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	dbUser , err := a.server.queries.GetUserByEmail(context.Background() , user.Email)

	if err == sql.ErrNoRows{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Incorrect email or passiword"})
		return
	} else if err != nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
        return
	}
	if err := utils.VerifyPassword(user.Password , dbUser.HashedPassword); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "incorrect email or password"})
		return
	}

	token , err := utils.CreateToken(dbUser.ID ,  a.server.config.Signed_key)

	
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()}) 
		return
	}
	c.JSON(http.StatusOK , gin.H{"token" : token})
}