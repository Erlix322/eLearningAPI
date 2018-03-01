package tokenhandler

import (
	"net/http"	
	"encoding/json"
	"fmt"
	"io/ioutil"	
	"time"
	"crypto/tls")

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
