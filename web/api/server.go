package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "simpleBank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("valid_currency", validCurrency)
		if err != nil {
			return nil
		}
	}

	//account
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.getListAccount)

	//transfer
	router.POST("/transfer", server.createTransfer)

	//user
	router.POST("/user", server.createUser)

	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func errorResponseStr(msg string) gin.H {
	return gin.H{"error": msg}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
