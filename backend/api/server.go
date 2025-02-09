package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/demmynile/fingo/db/sqlc"
	"github.com/demmynile/fingo/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router *gin.Engine
}


func NewServer(envPath string) *Server{

	config , err := utils.LoadConfig(envPath)

	if err != nil {
		panic(fmt.Sprintf("Could not load env config: %v" , err ))
	}

	conn , err := sql.Open(config.DBdriver, config.DB_source)
	if err != nil{
		panic("Could not connect to database")
	}

	q := db.New(conn)

	g := gin.Default()

	return &Server{
		queries: q,
		router: g,
	}
	
}


func (s *Server) Start(port int){
	s.router.GET("/" , func(ctx *gin.Context){
		ctx.JSON(http.StatusOK , gin.H{"message" : "Welcome to Fingreat"})
	})

	s.router.Run(fmt.Sprintf(":%v" , port))
}


