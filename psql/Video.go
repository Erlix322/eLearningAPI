package psql

type Video struct{
	Id int
	Name string `gorm:"column:Name"`
	Owner string `gorm:"column:Owner"`
	Beschreibung string `gorm:"column:Owner"`
}

type Videos struct{
	videos []Video
}

func (Video) TableName() string{
	return "Video"
}