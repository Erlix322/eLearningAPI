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
	fmt.Printf("WTF: %+v\n", pl.VideoPlaylist[0].PKVideo)
	err = db.Create(&pl.Playlist).Error
    /*create VideoPlaylist*/
	for index,playlist := range pl.VideoPlaylist{
		fmt.Println(playlist)
		play := pogo.VideoPlaylist{PKVideo:pl.VideoPlaylist[index].PKVideo,PKPlaylist:pl.VideoPlaylist[index].PKPlaylist}
		err = db.Table("VideoPlaylist").Create(&play).Error
	}
	if err != nil{
		tx.Rollback()
	}else{
		tx.Commit()
	}
	

}



