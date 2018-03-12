package psql

import (
	"log"
	"eLearningAPI/pogo"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)


func (c *Connection) SavePlaylist(pl pogo.VideoPlaylists){
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	/*create playlist*/
	db.Create(&pl.Playlist)
    /*create VideoPlaylist*/
	for playelem := range pl.List{
		db.Create(&playelem)
	}
	tx.Commit()

}



