package dbconn

import (
	"fmt"
	_ "github.com/apache/calcite-avatica-go/v5"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)


// singleton instance of database connection.
var dbInstance *gorm.DB
var dbOnce sync.Once

// DBPortalBRIBrain creates a new instance of gorm.DB if a connection is not established.
// return singleton instance.
func DB() *gorm.DB {
	if dbInstance == nil {
		dbOnce.Do(openDB)
	}
	return dbInstance
}

// openDB initialize gorm DBPortalBRIBrain.
func openDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_SCHEMA")) //Build connection string

	//fmt.Println(connectionString)
	file, err := os.Create("gorm-log.txt")
	if err != nil {
		// Handle error
		panic(err)
	}
	// Make sure file is closed before your app shuts down.
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	gormDB, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{SkipDefaultTransaction: true, Logger: newLogger},
	)

	if err != nil {
		panic("dbconn openDB PortalBRIBrain: cannot open database")
	}
	dbInstance = gormDB
	db, err := dbInstance.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}


