package migrations

import (
	"golang-clean-architecture/internal/entity"
	"log"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	var err error

	// auto migrate
	err = DB.AutoMigrate(&entity.User{}, &entity.Role{},
		&entity.Course{}, &entity.Section{}, &entity.Lesson{},
		&entity.Category{}, &entity.Review{}, &entity.Enrollment{},
		&entity.Transaction{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated successfully")
}
