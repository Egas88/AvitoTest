package domain

import "fmt"

type Segment struct {
	Id    int    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name  string `json:"name" gorm:"unique, not null"`
	Users []User `json:"users" gorm:"many2many:user_segment"`
}

func (segment *Segment) TableName() string {
	return "segment"
}

func (segment Segment) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", segment.Id, segment.Name)
}
