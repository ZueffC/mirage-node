package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db, err = gorm.Open("sqlite3", "test_database.db")

func Connect() {
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&PackageModel{})
}

func Database() *gorm.DB {
	return db
}
