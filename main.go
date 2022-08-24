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

	var pointName []*string
	names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	for _, name := range names {
		pointName = append(pointName, &name)
	}
	
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

