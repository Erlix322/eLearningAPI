package psql

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)


type Connection struct {
	connsTr string	
}

func NewConnection(connectionString string) *Connection{
	p:= &Connection{connsTr:connectionString}
	return p
}

func (c *Connection) connect(){
	db, err := sql.Open("postgres",c.connsTr)
	if err != nil {
		log.Fatal(err)
	}
}



