package main

import (
	"fmt"
	"interfaceAssigment/service"
)

func main() {
	var db []service.User

	res := service.NewUserService(db)
	fmt.Println(res)
	add := res.Register(&service.User{Nama: "Doni"})
	fmt.Println(add)
	get := res.Get()
	for _, each := range get {
		fmt.Println(each)
	}
}