package game

import (
	"log"
	"net/http"
)

func InitService() {

	http.HandleFunc("/user/login", UserLogin)
	http.HandleFunc("/user/add", AddUser)
	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}