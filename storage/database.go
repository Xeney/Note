package storage

import (
	"gin-notes-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Объявление глобальной переменной для хранения экземпляра базы данных
var db *gorm.DB

func InitDatabase() error {
	var err error
	// Открытие подключения к базе данных
	db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// Автоматическое создание таблицы для модели Note, если её ещё нет
	return db.AutoMigrate(&models.Note{})
}

// Функция для получения экземпляра базы данных
func GetDB() *gorm.DB {
	return db
}
