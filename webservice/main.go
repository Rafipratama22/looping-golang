package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webservice/service"
)

var PORT = ":8080"
var db []service.User
var userService = service.NewUserService(db)
var user service.User

func main() {
	http.HandleFunc("/", greetings)
	http.HandleFunc("/user", Get)
	http.HandleFunc("/register", Register)
	
	
	fmt.Println("server is listen in port ", PORT)
	http.ListenAndServe(PORT, nil)

}

func greetings(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello world")
}

func Register(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	fmt.Println(user)
	userService.Register(&user)
	fmt.Fprintf(w, "sukses to create")
}

func Get(w http.ResponseWriter, r *http.Request) {
	allUser := userService.Get()
	data, err := json.Marshal(allUser)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	json.NewEncoder(w).Encode(allUser)
}

func Update(w http.ResponseWriter, r *http.Request) {
	
}
