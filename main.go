package main

import (
	"net/http"	
	"strings"
	"encoding/json"
	"github.com/gorilla/mux" 
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"crypto/tls"
	"eLearningAPI/tokenhandler"
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



func Middleware(h http.Handler) http.Handler {
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
			h.ServeHTTP(res,req)
		}		
		
	})
}

func main() {
    //fs := http.FileServer(http.Dir("."))
    //http.Handle("/", http.StripPrefix("/", fs))
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/vid/", HomeHandler)
	r.HandleFunc("/vid/{key}", serveVideo).Methods("GET")
	//http.Handle("/",HomeHandler)
	http.Handle("/", r)
	//Secure Route
	http.Handle("/vid/", Middleware(r))
	http.ListenAndServe(":3000",nil)
}