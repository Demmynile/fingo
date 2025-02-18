package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/demmynile/fingo/db/sqlc"
	"github.com/demmynile/fingo/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	_ "github.com/lib/pq"
)

type Server struct {
	queries *db.Queries
	router *gin.Engine
	config *utils.Config
	
}

var tokenController *utils.JWTToken

var gValid = galidator.New().CustomMessages(
	galidator.Messages{
		"required": "this field is required",
	},
)

func myCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	return cors.New(config)
}



func NewServer(envPath string) *Server{

	config , err := utils.LoadConfig(envPath)

	if err != nil {
		panic(fmt.Sprintf("Could not load env config: %v" , err ))
	}

	conn , err := sql.Open(config.DBdriver, config.DB_source_live)
	fmt.Println("error:", err)

	if err != nil{
		panic("Could not connect to database")
	}

	tokenController = utils.NewJWTToken(config)

	q := db.New(conn)

	g := gin.Default()

	g.Use(myCorsHandler())

	return &Server{
		queries: q,
		router: g,
		config: config,
		
	}
	

}


func (s *Server) Start(port int){
	s.router.GET("/" , func(ctx *gin.Context){
		ctx.JSON(http.StatusOK , gin.H{"message" : "Welcome to Fingreat"})
	})

	User{}.router(s)
	Auth{}.router(s)



	s.router.Run(fmt.Sprintf(":%v" , port))
}


