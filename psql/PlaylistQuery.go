package psql

import (
	"log"
	"fmt"
	"eLearningAPI/pogo"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)


func (c *Connection) SavePlaylist(pl pogo.VideoPlaylists) bool{
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	/*create playlist*/
	if err := db.Create(&pl.Playlist).Error; err != nil {
		tx.Rollback()
		return false
	}
	/*create VideoPlaylist*/
	
	for index,playlist := range pl.VideoPlaylist{
		fmt.Println(playlist)
		play := pogo.VideoPlaylist{PKVideo:pl.VideoPlaylist[index].PKVideo,PKPlaylist:pl.VideoPlaylist[index].PKPlaylist}
		if err := db.Table("VideoPlaylist").Create(&play).Error; err != nil{
			tx.Rollback()
			return false
		}
	}
	
	tx.Commit()
	return true
	

}



