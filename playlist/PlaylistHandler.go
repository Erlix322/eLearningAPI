package playlist

import (
	"net/http"
	"encoding/json"
	"eLearningAPI/psql"
	"eLearningAPI/pogo"
	"fmt"	
)

func SavePlaylist(w http.ResponseWriter, r *http.Request) {
	
	con := psql.NewConnectionP()
	
	decoder := json.NewDecoder(r.Body)

	var t pogo.VideoPlaylists
	err := decoder.Decode(&t)

	if err != nil {
		fmt.Println(err)
	}
	/*Save playlist*/
	con.SavePlaylist(t)
	
}