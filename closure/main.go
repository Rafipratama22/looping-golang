package closure

import "fmt"

func main() {
	name1 := "fisaa"
	name2 := "david"
	name3 := "fadly"

	pointName := []*string{&name1, &name2, &name3}
	funName := func(pointName []*string) {
		for _, point := range pointName {
			fmt.Println(*point)
		}
	}
	funName(pointName)
}
