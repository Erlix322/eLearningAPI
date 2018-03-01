package settingshandler
import (
	"net/http"		
	"time"
	"fmt"
	"io/ioutil"
	"crypto/tls"
)
type SettingsHandler struct{}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport:tr,
	}
	req, err := http.NewRequest("GET","https://rubu2.rz.htw-dresden.de/API/v0/studentTimetable.php?StgJhr=17&Stg=044&StgGrp=72",nil)
	if err != nil {
		fmt.Println(err)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
    bodyString := string(body)
	fmt.Fprintf(w, bodyString)
}



