package psql

import (
	_ "github.com/go-sql-driver/mysql"
)

type Video struct{
	Id int `gorm:"column:Id"`
	Name string `gorm:"column:Name"`
	Owner string `gorm:"column:Owner"`
	Beschreibung string `gorm:"column:Beschreibung"`
}

type Videos struct{
	videos []Video
}

func (Video) TableName() string{
	return "Video"
}