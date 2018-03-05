package psql

import (
	"database/sql"
	"log"
	"fmt"
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
	stmt, err := db.Prepare("INSERT INTO Video(Name) VALUES(?)")
	
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(video)
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
	db, err := sql.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id,Name,Modul,Beschreibung from Video;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	vs := Videos{}
	for rows.Next() {
		v := Video{}
		err := rows.Scan(&v.Id,&v.Name,&v.Modul,&v.Beschreibung)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(v.Id, v.Name)
		vs.videos = append(vs.videos,v)		
	}

	/*TEST*/
	/*	vs := Videos{}
		v := Video{Id:1,Name:"hallo"}
		vs.videos = append(vs.videos,v)
	*/

	return vs.videos
}



