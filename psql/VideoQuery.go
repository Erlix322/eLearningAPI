package psql

import (
	"database/sql"
	"log"
	"fmt"
	"os"
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

func NewConnectionP() *Connection{
	var user = os.Args[1]
	var password = os.Args[2]
	var database = os.Args[3]
	connectionString :=""+user+":"+password+"@/"+database+""
	p:= &Connection{connsTr:connectionString}
	return p
}

func (c *Connection) SaveVideo(video string,owner string) int64{
	db, err := sql.Open("mysql",c.connsTr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO Video(Name,Modul,Owner,Beschreibung) VALUES(?,?,?,?)")
	
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(video,"",owner,"")
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
	
	db.Raw("Select Id,Name,Owner,Beschreibung from Video;").Find(&vs.videos)
	fmt.Println(vs.videos)
	return vs.videos
}

func (c *Connection) GetVideosByUser(user string) []Video{
	db, err := gorm.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
	vs := Videos{}
	
	db.Where("Owner = ?",user).Find(&vs.videos)
	return vs.videos
}





