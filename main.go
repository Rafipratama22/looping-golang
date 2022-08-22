package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			fmt.Println(i, "ganjil")
		} else {
			fmt.Println(i, "genap")
		}
	}

	names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	for _, n := range names {
		fmt.Println(n)
	}
}