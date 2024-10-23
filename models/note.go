package models

type Note struct {
	// Уникальный идентификатор заметки, автоматически инкрементируемый
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	// Заголовок заметки
	Title string `json:"title"`
	// Содержание заметки
	Content string `json:"content"`
}
