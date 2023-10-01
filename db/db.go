package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yuki0418/go-gin-tutorial/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseManager struct {
}

func  (d *DatabaseManager) Init() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_DATABASE")
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	// print(dsn)
	println(dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, 
	}), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
	)
}

func (d *DatabaseManager) GetTestDb(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	return db, mock
}