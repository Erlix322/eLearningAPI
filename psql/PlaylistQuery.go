package psql

import (
	"log"
	"fmt"
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
	err = db.Create(&pl.Playlist).Error
    /*create VideoPlaylist*/
	for playelem := range pl.VideoPlaylist{
		fmt.Println(playelem)
		err = db.Create(&playelem).Error
	}
	if err != nil{
		tx.Rollback()
	}else{
		tx.Commit()
	}
	

}



