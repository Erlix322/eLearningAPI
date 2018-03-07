package psql

type VideoPlaylist struct {
	PKVideo int
	PKPlaylist int
}

func (VideoPlaylist) TableName() string{
	return "VideoPlaylist"
}