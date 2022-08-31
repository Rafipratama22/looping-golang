package main

import (
	"fmt"
	"interfaceAssigment/service"
	"sync"
)

func main() {
	var db []service.User
	var wg sync.WaitGroup

	res := service.NewUserService(db)
	fmt.Println(res)
	wg.Add(3)
	go res.Register(&service.User{Nama: "Doni"}, &wg)

	go res.Register(&service.User{Nama: "David"}, &wg)
	
	go res.Get(&wg)
	wg.Wait()
}