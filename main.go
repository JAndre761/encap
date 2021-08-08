package main

import (
	"fmt"
	"github.com/JAndre761/encap/course"
)

func main() {
	// Go es una variable, se le asigna un objeto de instancia de Course
	Go := course.New("Go from nothing", 12.4, false)

	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	// fmt.Printf("%+v", Go)

}
