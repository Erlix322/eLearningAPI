package main

import (
	"net/http"	
	"strings"
	"github.com/rs/cors"
	"github.com/gorilla/mux" 
	"fmt"
	"os"
	"time"
	"eLearningAPI/tokenhandler"
	"eLearningAPI/settingshandler"
	"eLearningAPI/session"
)

func serveVideo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key,ok := vars["key"]
	fmt.Println("Video is",ok,key)	
	video, err := os.Open("./"+key+".mp4")
	defer video.Close()
	if err != nil {
		ErrorHandler(w,r)
	}else {
		http.ServeContent(w,r, key+".mp4",time.Now(),video)
	}
}

func HomeHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Hello home")
}

func ErrorHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Failed Authentication")
}


/*
Check if Authorization Token is set and create Session if it is present
*/
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		var token string
		tokens, ok := req.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}
		if token == "" {
            // If we get here, the required token is missing
            http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
            return
		}
		
		th := tokenhandler.NewTokenHandler(token)
		
		ret,val := th.CheckToken(token)
		fmt.Println(val)
		if(ret){
			s.SetSession(res,req,"video")
			h.ServeHTTP(res,req)
		}		
		
	})
}

func SessionMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		if(s.CheckSession(res,req,"video")){
			h.ServeHTTP(res,req)
		}		
	})
}
var s = session.NewSession()

func main() {
    //fs := http.FileServer(http.Dir("."))
	//http.Handle("/", http.StripPrefix("/", fs))
	s.CreateCookieStore("superSecret")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/vid/", HomeHandler)
	r.HandleFunc("/vid/{key}", serveVideo).Methods("GET")
	r.HandleFunc("/settings", settingshandler.GetSettings )
	//http.Handle("/",HomeHandler)
	//corsObj:=r.AllowedOrigins([]string{"*"})
	

	r2 := mux.NewRouter()
	r2.HandleFunc("/auth/", HomeHandler)
	r2.HandleFunc("/auth/vid/",serveVideo)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
	
	//Secure Route
	http.Handle("/", c.Handler(r))
	http.Handle("/auth/",c.Handler(AuthMiddleware(r)))
	http.Handle("/vid/", c.Handler(SessionMiddleWare(r)))
	http.ListenAndServe(":3001",nil)
}