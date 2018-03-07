package psql

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)


type Connection struct {
	connsTr string
}

func NewConnection(connectionString string) *Connection{
	p:= &Connection{connsTr:connectionString}
	return p
}

func (c *Connection) SaveVideo(video string) int64{
	db, err := sql.Open("mysql",c.connsTr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO Video(Name,Modul,Beschreibung,Owner) VALUES(?,?,?,?)")
	
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(video,"","","")
	if err != nil {
		fmt.Println(err)
	}
	id,err := result.LastInsertId()
	fmt.Println("id:",id)
	if err != nil {
		fmt.Println(err)
	}	
	return id

}

func (c *Connection) GetVideos() []Video{
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	vs := Videos{}
	
	db.Find(&vs.videos)
	return vs.videos
}

func (c *Connection) SavePlaylist(pl []VideoPlaylist){
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	for playelem := range pl{
		db.Create(&playelem)
	}
	tx.Commit()

}



