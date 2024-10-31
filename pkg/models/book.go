package models

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Content []Book `json:"content"`
}
