package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	db "github.com/justIGreK/simple_bank_course/db/sqlc"
	"github.com/justIGreK/simple_bank_course/db/token"
	"github.com/justIGreK/simple_bank_course/db/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter(){
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	
	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
