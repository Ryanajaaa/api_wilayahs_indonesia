package models

type Province struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `gorm:"type:varchar(50);unique" json:"code"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}
