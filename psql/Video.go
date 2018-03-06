package psql

type Video struct{
	Id int
	Name string
	Owner string
	Beschreibung string
}

type Videos struct{
	videos []Video
}

func (Video) TableName() string{
	return "Video"
}