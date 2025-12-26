package bootstrap

import (
	"cash-flow/server/internal/config"
	"cash-flow/server/internal/database"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Mongo  *mongo.Database
}

func NewApp() *App {
	cfg := config.Load()

	client, err := database.NewMongoClient(cfg.Mongo.URI)
	if err != nil {
		log.Fatal("mongo connection failed:", err)
	}

	db := client.Database(cfg.Mongo.Database)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	return &App{
		Router: router,
		Mongo:  db,
	}
}

func (a *App) Run() {
	a.Router.Run(":" + config.Load().Server.Port)
}
