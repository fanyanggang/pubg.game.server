package game

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func AddUser(response http.ResponseWriter, request *http.Request) {
}

func UserLogin(response http.ResponseWriter, request *http.Request) {

	data := map[string]string{
		"test": "hello",
	}

	fmt.Println(request.URL.Path) 
	temp,_ := json.Marshal(data)
	fmt.Fprintf(response, string(temp))

	fmt.Print("test")
}

func AddGroup(response http.ResponseWriter, request *http.Request) {
}

