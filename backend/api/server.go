package api

import (
	"database/sql"
	db "github/demmynile/fingo/db/sqlc"
	"github/demmynile/fingo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router *gin.Engine
}


func NewServer(envPath string) *Server{

	config , err := utils.LoadConfig("../..")

	if err != nil {
		panic("Could not load env config")
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
}


