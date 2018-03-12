package pogo

type Playlist struct{
	Id int `gorm:"column:id"`
	Name string `gorm:"column:Name"`
	Beschreibung string `gorm:"column:Beschreibung"`
}

func (Playlist) TableName() string{
	return "Playlist"
}