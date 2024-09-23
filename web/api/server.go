package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "simpleBank/db/sqlc"
	"simpleBank/token"
	"simpleBank/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.TokenMaker

	store  *db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {

	//maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	maker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: maker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("valid_currency", validCurrency)
		if err != nil {
			return nil, nil
		}
	}

	//account
	server.setupRouter(router)

	return server, nil
}

func (server *Server) setupRouter(router *gin.Engine) {

	//login
	router.POST("/login", server.loginUser)

	authRouters := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//user
	authRouters.POST("/user", server.createUser)

	//account
	authRouters.POST("/account", server.createAccount)
	authRouters.GET("/account/:id", server.getAccount)
	authRouters.GET("/account", server.getListAccount)

	//transfer
	authRouters.POST("/transfer", server.createTransfer)

	server.router = router
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
