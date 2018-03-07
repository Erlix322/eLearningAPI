package tokenhandler

import (
	"net/http"	
	"encoding/json"
	"fmt"
	"io/ioutil"	
	"time"
	"crypto/tls"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

type TokenHandler struct {
	token string
}

func NewTokenHandler(token string) *TokenHandler{
	p:= &TokenHandler{token:token}
	return p
}

func (h *TokenHandler) CheckToken(token string) (bool, string){
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport:tr,
	}
	req, err := http.NewRequest("GET","https://api.brandt-projects.de/auth/realms/eLearning/protocol/openid-connect/userinfo",nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
    bodyString := string(body)
	fmt.Println("String",bodyString)
	defer resp.Body.Close()
	
	c := make(map[string]interface{})
	e := json.Unmarshal([]byte(bodyString),&c)
	if e != nil {
		fmt.Println("Error parsing string")
	}

	if k, ok := c["sub"]; ok{
		return true,k.(string)
	}

	return false, "invalid token"
	

}

func GetUsername(req *http.Request) string{
	var tokenString string
	tokens, ok := req.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		tokenString = tokens[0]
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return nil,nil
	})
	
	fmt.Println(token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok  {
		return claims["preferred_username"].(string)
	} else {
		return ""
	}
}
