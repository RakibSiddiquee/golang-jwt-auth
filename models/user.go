package models

// json is used to return column name
// json:"-" means password will not return with the response
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password []byte `json:"-"`
}
