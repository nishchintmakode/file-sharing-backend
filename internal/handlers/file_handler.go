package handlers

import (
	"file-sharing-backend/internal/models"
	"file-sharing-backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileHandler struct {
	db       *gorm.DB
	s3Client *storage.S3Client
	cache    *utils.RedisClient
}

func NewFileHandler(r *gin.RouterGroup, db *gorm.DB, s3Client *storage.S3Client, cache *utils.RedisClient) {
	h := &FileHandler{db: db, s3Client: s3Client, cache: cache}
	r.POST("/upload", h.UploadFile)
	r.GET("/files", h.ListFiles)
	r.GET("/share/:file_id", h.ShareFile)
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	s3URL, err := h.s3Client.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed"})
		return
	}

	newFile := models.File{
		UserID:   userID,
		Filename: file.Filename,
		Size:     file.Size,
		S3URL:    s3URL,
	}
	h.db.Create(&newFile)
	h.cache.Delete(c, "user_files_"+strconv.Itoa(int(userID))) // Invalidate cache
	c.JSON(http.StatusOK, gin.H{"url": s3URL})
}

func (h *FileHandler) ListFiles(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var files []models.File

	// Check cache first
	cacheKey := "user_files_" + strconv.Itoa(int(userID))
	if cached := h.cache.Get(c, cacheKey); cached != nil {
		c.JSON(http.StatusOK, cached)
		return
	}

	// Query DB if cache miss
	h.db.Where("user_id = ?", userID).Find(&files)
	h.cache.Set(c, cacheKey, files, 300) // Cache for 5 minutes
	c.JSON(http.StatusOK, files)
}
