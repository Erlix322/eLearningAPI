package psql

type Video struct{
	Id int
	Name string
	Modul string
	Beschreibung string
}

type Videos struct{
	videos []Video
}