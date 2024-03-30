package models

// Student represents the structure of a student entity.
type Student struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}
