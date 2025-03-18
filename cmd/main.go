package main

import (
	"file-sharing-backend/internal/auth"
	"file-sharing-backend/internal/handlers"
	"file-sharing-backend/internal/workers"
	"file-sharing-backend/pkg/config"
	"file-sharing-backend/pkg/utils"
	"file-sharing-backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := utils.InitDB(cfg)
	defer utils.CloseDB(db)
	s3Client := storage.NewS3Client(cfg)
	redisClient := utils.InitRedis(cfg)

	// Start background worker
	go workers.StartCleanupWorker(db, s3Client)

	r := gin.Default()
	authMiddleware := auth.AuthMiddleware(cfg)
	api := r.Group("/api")
	{
		handlers.NewFileHandler(api, db, s3Client, redisClient)
	}
	r.Run(":" + cfg.AppPort)
}
