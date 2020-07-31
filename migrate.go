
package main

import (
	"simpleapi/configs"
	"simpleapi/models"
)

func main() {
	db := configs.GetDBConnection()
	db.AutoMigrate(&models.Message{})	
}
