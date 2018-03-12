package psql

import (
	"log"
	"fmt"
	"eLearningAPI/pogo"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	_ "github.com/go-sql-driver/mysql"
)


func (c *Connection) SavePlaylist(pl pogo.VideoPlaylists) bool{

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return false
	}
	uuid := u2.String()
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	/* create playlist */
	pl.Playlist.Id = uuid
	if err := db.Create(&pl.Playlist).Error; err != nil {
		tx.Rollback()
		return false
	}
	/*create VideoPlaylist*/
	
	for index,playlist := range pl.VideoPlaylist{
		fmt.Println(playlist)
		play := pogo.VideoPlaylist{PKVideo:pl.VideoPlaylist[index].PKVideo,PKPlaylist:uuid}
		if err := db.Table("VideoPlaylist").Create(&play).Error; err != nil{
			tx.Rollback()
			return false
		}
	}
	
	tx.Commit()
	return true
	

}



