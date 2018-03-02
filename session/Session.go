package session

import(
	"github.com/gorilla/sessions"
	"net/http"
)

type Session struct{
	store *sessions.CookieStore
}

func NewSession() *Session{
	p:= &Session{}
	return p
}


func (s *Session) CreateCookieStore(secret string){
	s.store = sessions.NewCookieStore([]byte(secret))    
}

func (s *Session) SetSession(w http.ResponseWriter, r *http.Request, clientID string) {
	session, err := s.store.Get(r,clientID)
	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	session.Values["legit"] = true
	session.Save(r,w)
}

func (s *Session) CheckSession(w http.ResponseWriter, r *http.Request, clientID string) bool{
	session, err := s.store.Get(r,clientID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if session.Values["legit"] == true	{
		return true
	}
	return false
}

func (s *Session) ClearSession(w http.ResponseWriter, r *http.Request, clientID string) {
	session, err := s.store.Get(r,clientID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Options.MaxAge = -1
	session.Save(r,w)
}


