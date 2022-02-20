package config

import (
	"fmt"
	"log"
	"project-4/model"
	"project-4/tool"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// local
var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbname   = "hacktiv8-project-4"
	db       *gorm.DB
	err      error
)

// var (
// 	host     = os.Getenv("DB_HOST")
// 	user     = os.Getenv("DB_USER")
// 	password = os.Getenv("DB_PASSWORD")
// 	dbPort   = os.Getenv("DB_PORT")
// 	dbname   = os.Getenv("DB_NAME")
// 	db       *gorm.DB
// 	err      error
// )

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: InitLog(),
	})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	if err = db.AutoMigrate(model.Category{}, model.Product{}, model.TransactionHistory{}); err != nil {
		log.Fatal("Error run database automigration", err)
	}

	if err = db.AutoMigrate(model.User{}); err == nil && db.Migrator().HasTable(model.User{}) {
		var count int64
		db.Model(&model.User{}).Count(&count)

		if count == 0 {
			admin := model.User{Fullname: "admin", Role: "admin", Password: tool.HashPassword("adminadmin"), Email: "admin@gmail.com", CreatedAt: time.Now(), Balance: 0}

			if err = db.Model(&model.User{}).Create(&admin).Error; err != nil {
				log.Fatal("Error create admin")
			}
		}
	} else {
		log.Fatal("Error run database automigration:", err)
	}

	fmt.Println("Database connection success.")
}

func GetDB() *gorm.DB {
	return db
}
