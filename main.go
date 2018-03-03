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
	"eLearningAPI/psql"
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
	conn := psql.NewConnection("user:password@/dbname")
	vids := conn.GetVideos()
	fmt.Println(vids)
}
/*
func verifyToken(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	fmt.Println(vars["token"])

	token, _ := jwt.Parse(vars["token"], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		//hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("kGxQIQIDAQAB"), nil
	})

	fmt.Println(token)
	
}
*/
func ErrorHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Video not found")
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
		}else{
			s.ClearSession(res,req,"video")
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
	//r.HandleFunc("/token/{token}",verifyToken)
	//http.Handle("/",HomeHandler)
	//corsObj:=r.AllowedOrigins([]string{"*"})

	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
	
	//Secure Route
	http.Handle("/", c.Handler(r))
	//http.Handle("/token/{token}", r)
	http.Handle("/auth/",c.Handler(AuthMiddleware(r)))
	http.Handle("/vid/", c.Handler(SessionMiddleWare(r)))
	http.ListenAndServe(":3001",nil)
}