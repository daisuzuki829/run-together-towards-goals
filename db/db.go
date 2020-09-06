package db

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
)

// dbInit...
func Init() *gorm.DB {
	db := models.Open()
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(
		&models.User{},
		&models.DailyKpt{},
		&models.Goal{},
		&models.Genre{},
		&models.KptReactionHistory{},
		&models.TodoList{},
		&models.MonthlyPlan{},
	)

	r := models.NewGenreRepository()
	r.GenreMigration()

	return db
}
