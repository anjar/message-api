package configs

import (
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

type dbUtil struct {
	db *gorm.DB
}

var dbInstance *dbUtil
var dbOnce sync.Once

// gets Database (MySQL) connection
func GetDBConnection() *gorm.DB {
	dbOnce.Do(func() {
		log.Println("Initialize DB connection...")
		conn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(os.Getenv("DB_TYPE"), conn)

		if err != nil {
			panic(err)
		}

		/**
		 * NOTES: this will set connection lifetime in connection pool to 1 minute.
		 * 		  If the connection in the pool is idle > 1 min, Golang will close it
		 * 		  and will create new connection if #connections in the pool < pool max num
		 * 		  of connection. This to avoid invalid connection issue
		 */
		db.DB().SetConnMaxLifetime(time.Second * 60)
		db.SingularTable(true) // Set as singular table
		db.LogMode(false)

		if err != nil {
			panic(err)
		}

		dbInstance = &dbUtil{
			db: db,
		}
	})

	return dbInstance.db
}
