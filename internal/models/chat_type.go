package models

type ChatType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
