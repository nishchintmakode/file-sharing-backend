package workers

import (
	"file-sharing-backend/internal/models"
	"file-sharing-backend/internal/storage"
	"time"

	"gorm.io/gorm"
)

func StartCleanupWorker(db *gorm.DB, s3Client *storage.S3Client) {
	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		var expiredFiles []models.File
		db.Where("upload_date < ?", time.Now().Add(-7*24*time.Hour)).Find(&expiredFiles)
		for _, file := range expiredFiles {
			s3Client.DeleteFile(file.S3URL) // Implement S3 delete
			db.Delete(&file)
		}
	}
}
