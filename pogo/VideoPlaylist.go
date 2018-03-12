package pogo
type VideoPlaylist struct{
	PKVideo int  `gorm:"column:PKVideo"`
	PKPlaylist int `gorm:"column:PKPlaylist"`
}

func (VideoPlaylist) TableName() string{
	return "VideoPlaylist"
}

type VideoPlaylists struct{
	Playlist Playlist `json:"playlist"`
	VideoPlaylist []VideoPlaylist `json:"list"`
}


/*
JSON Format:
{
   "playlist":{
      "id":1,
      "Name":"huhu",
      "Beschreibung":"test"
   },
   "list":[
      {
         "PKVideo":1,
         "PKPlaylist":2
	  },
	  {
		 "PKVideo":2,
         "PKPlaylist":2
	  }
   ]
}
*/