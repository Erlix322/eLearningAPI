package uploader

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"mime/multipart"
	"strconv"
    "eLearningAPI/psql"
    "eLearningAPI/tokenhandler"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPload gestartet")
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    file, handle, err := r.FormFile("file")
    if err != nil {
        fmt.Fprintf(w, "%v", err)
        return
    }
    defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	fmt.Println(mimeType)
    switch mimeType {
    case "video/mp4":
        saveFile(w,r, file, handle)
    default:
        jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
    }
}

func saveFile(w http.ResponseWriter, r *http.Request, file multipart.File, handle *multipart.FileHeader) {
	var user = os.Args[1]
	var password = os.Args[2]
	var database = os.Args[3]
	fmt.Println("user:",user)
	conn := psql.NewConnection(""+user+":"+password+"@/"+database+"")
    fmt.Println("Connected")
    owner := tokenhandler.GetUsername(r)
	id := conn.SaveVideo(handle.Filename,owner)
	fmt.Println(id)
	data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Fprintf(w, "%v", err)
        return
    }
    err = ioutil.WriteFile("./files/"+strconv.Itoa(int(id))+".mp4", data, 0666)
    if err != nil {
        fmt.Fprintf(w, "%v", err)
        return
	}	

    jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    fmt.Fprint(w, message)
}