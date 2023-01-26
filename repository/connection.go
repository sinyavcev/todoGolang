package repository

import (
	"github.com/sinyavcev/todoGolang/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() (*gorm.DB, error) {

	dsn := "host=127.0.0.1 user=djamware password=dj@mw@r3 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Todo{})

	return db, nil
}
