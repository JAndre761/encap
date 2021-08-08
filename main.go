package main

func main() {
	// Go es una variable, se le asigna un objeto de instancia de Course
	Go := Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()

}
