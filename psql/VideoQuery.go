package psql

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


type Connection struct {
	connsTr string
}

func NewConnection(connectionString string) *Connection{
	p:= &Connection{connsTr:connectionString}
	return p
}

func (c *Connection) GetVideos() Videos{
	db, err := sql.Open("mysql",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id,Name from Video;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	vs := Videos{}
	for rows.Next() {
		v := Video{}
		err := rows.Scan(v.Id,v.Name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(v.Id, v.Name)
		vs.videos = append(vs.videos,v)		
	}

	return vs
}



