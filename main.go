package main

import (
	"net/http"	
	"strings"
	"github.com/rs/cors"
	"github.com/gorilla/mux" 
	//"github.com/gorilla/handlers"
	"fmt"
	"os"
	
	"time"	
	"encoding/json"
	"eLearningAPI/tokenhandler"
	"eLearningAPI/settingshandler"
	"eLearningAPI/uploader"
	"eLearningAPI/psql"
)

func serveVideo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key,ok := vars["key"]
	fmt.Println("Video is",ok,key)	
	video, err := os.Open("./files/"+key+".mp4")
	defer video.Close()
	if err != nil {
		ErrorHandler(w,r)
	}else {
		http.ServeContent(w,r, key+".mp4",time.Now(),video)
	}
}

func HomeHandler(res http.ResponseWriter, req *http.Request){
	var user = os.Args[1]
	var password = os.Args[2]
	var database = os.Args[3]
	conn := psql.NewConnection(""+user+":"+password+"@/"+database+"")
	owner := tokenhandler.GetUsername(r)
	vids := conn.GetVideosByUser(owner)
	m,_:=json.Marshal(vids)
	fmt.Fprintf(res,string(m))
}

func ErrorHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Video not found")
}


/*
Check if Authorization Token is set and create Session if it is present
*/
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		//res.Header().set("Access-Control-Allow-Origin","*")
		var token string
		tokens, ok := req.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}
		if token == "" {
            // If we get here, the required token is missing
            http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
            return
		}
		
		th := tokenhandler.NewTokenHandler(token)
		
		ret,val := th.CheckToken(token)
		fmt.Println(val)
		if(ret){								
			h.ServeHTTP(w,req)
		}		
	})
}


func SessionMiddleWare(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		vals := req.URL.Query() // Returns a url.Values, which is a map[string][]string
		token, ok := vals["access_token"] // Note type, not ID. ID wasn't specified anywhere.
		
		if ok {
			if len(token) >= 1 {
				th := tokenhandler.NewTokenHandler(token[0])
				ret,val := th.CheckToken(token[0])
				fmt.Println(val)
				if(ret){								
					h.ServeHTTP(w,req)
				}						
			}
		}
	})
}


func main() {
/*	ok := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin","X-Requested-With","Content-Type", "Access-Control-Allow-Headers", "Authorization"})
	ok2 := handlers.AllowedOrigins([]string{"*"})
	ok3 := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	*/
	
		
	
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth/", HomeHandler)
	r.HandleFunc("/vid/", HomeHandler)
	r.HandleFunc("/upload/",uploader.UploadFile)
	r.HandleFunc("/vid/{key}", serveVideo)
	r.HandleFunc("/settings", settingshandler.GetSettings )
	

	
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Authorization","Content-Type"},
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
	
	//Secure Route
	http.Handle("/", c.Handler(r))

	//http.Handle("/token/{token}", r)
	http.Handle("/upload/",c.Handler(AuthMiddleware(r)))
	http.Handle("/auth/",c.Handler(AuthMiddleware(r)))
	http.Handle("/vid/",c.Handler(SessionMiddleWare(r)))
	http.ListenAndServe(":3001",nil)
}

