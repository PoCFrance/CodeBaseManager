package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Server struct {
	Port   string
	Router *gin.Engine
	Db     *database.Database
}

func NewServer() (*Server, func()) {
	server := &Server{}
	server.Init()
	return server, server.Destroy
}

func (s *Server) Init() {
	s.Port = os.Getenv("PORT")
	if s.Port == "" {
		s.Port = "5342"
		log.Printf("Defaulting to port %s", s.Port)
	}
	var err error
	s.Db, err = database.Init()
	if err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}
	models.MigrateModels(s.Db)
	s.Router = gin.Default()
	s.Router.Use(database.SetDatabase(s.Db))
}

func (s *Server) Destroy() {
	s.Db.Destroy()
}
