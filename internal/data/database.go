package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})

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
