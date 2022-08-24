package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		fmt.Println(i, "ganjil")
	// 	} else {
	// 		fmt.Println(i, "genap")
	// 	}
	// }

	// names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	name1 := "fisaa"
	name2 := "david"
	name3 := "fadly"

	pointName := []*string{&name1, &name2, &name3}
	funName := func (pointName []*string)  {
		for _, point := range pointName {
			fmt.Println(*point)
		}	
	}
	funName(pointName)
	// callName(names...)
}

func callName(names ...string) {
	for _, n := range names {
		fmt.Println(n)
	}	
}

