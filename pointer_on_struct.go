package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	me := person{"Antonio", 19}
	changeMe(&me)
	fmt.Println(me)
}

func changeMe(p1 *person) {
	p1.name = "Antonio Carlos"
	// or (*p1).name = "Antonio Carlos"
}
