package closure

import "fmt"

func main() {
	var pointName []*string
	names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	for _, name := range names {
		pointName = append(pointName, &name)
	}
	// pointName := []*string{&name1, &name2, &name3, &name4, &name5}
	funName := func(pointName []*string) {
		for _, point := range pointName {
			fmt.Println(*point)
		}
	}
	funName(pointName)
}
