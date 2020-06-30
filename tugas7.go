package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var x int
	var y string

	go bacatipe2(y)
	bacatipe1(x)

	var input string
	fmt.Scanln(&input)
}

func bacatipe1(item int) {
	reflectVar := reflect.ValueOf(item)
	fmt.Println("Tipe Variabel: ", reflectVar.Type())
}

func bacatipe2(item string) {
	reflectVar := reflect.ValueOf(item)
	fmt.Println("Tipe Variabel: ", reflectVar.Type())
}
