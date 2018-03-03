package psql

type Video struct{
	Id int
	Name string
}

type Videos struct{
	videos []Video
}