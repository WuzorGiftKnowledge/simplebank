package api

import (
	db "github.com/WuzorGiftKnowledge/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/createaccount", server.createAccount)
	router.GET("/getaccount:id", server.getAccount)
	router.GET("/listaccounts", server.listAccounts)
	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
