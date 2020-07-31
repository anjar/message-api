package repositories

import (
	"simpleapi/configs"
	"github.com/jinzhu/gorm"
)
type BaseRepositories struct {
	Database *gorm.DB
}

func NewBaseRepositories() *BaseRepositories {
	// set database
	db := configs.GetDBConnection()
	// set base repositories
	return &BaseRepositories{
		Database: db,
	}
}