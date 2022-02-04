package  database
import (
	"gorm.io/gorm"
)

func SyncDatabaseTables(db *gorm.DB){
	db.AutoMigrate(&Project{})
}
