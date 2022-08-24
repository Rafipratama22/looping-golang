package main

import "fmt"

type nameStruct struct {
	named string
}

func main() {
	// var pointName []*string
	var pointName []*nameStruct
	names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	for _, name := range names {
		// pointName = append(pointName, &name)
		named := nameStruct{
			named: name,
		}
		pointName = append(pointName, &named)
	}
	// pointName := []*string{&name1, &name2, &name3, &name4, &name5}
	funName := func(pointName []*nameStruct) {
		for _, point := range pointName {
			fmt.Println(point.named)
		}
	}
	funName(pointName)
}
