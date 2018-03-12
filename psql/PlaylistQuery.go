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
	for videoplaylist := range pl.VideoPlaylist{
		fmt.Printf("WTF: %+v\n", videoplaylist)
		err = db.Table("VideoPlaylist").Create(&videoplaylist).Error
	}
	if err != nil{
		tx.Rollback()
	}else{
		tx.Commit()
	}
	

}



