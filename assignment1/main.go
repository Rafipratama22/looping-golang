package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	name string
	alamat string
	pekerjaan string
	reason string
}

func main() {
	// names := []string{"Fadli", "David", "Fisa", "Teguh", "Fajar", "Irfan", "Edwin", "Hans", "Jaka", "Wicaksana"}
	students := []student{
		{
			name: "Fadli",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "David",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Fisa",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Teguh",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Fajar",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Irfan",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Edwin",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Hans",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Jaka",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
		{
			name: "Wicaksana",
			alamat: "jakarta",
			pekerjaan: "developer",
			reason: "suka aja",
		},
	}
	args := os.Args
	indexArgs := args[1]
	index, _ := strconv.Atoi(indexArgs)
	for i, student := range students {
		if i == index - 1 {
			fmt.Println(student)
		}
	}
}