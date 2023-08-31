package domain

import "fmt"

type User struct {
	Id       int       `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name     string    `json:"name"`
	Segments []Segment `json:"segments" gorm:"many2many:user_segment"`
}

func (user *User) TableName() string {
	return "users"
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", user.Id, user.Name)
}
