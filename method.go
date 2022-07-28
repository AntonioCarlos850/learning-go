package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) getNameAndAge() (string, int) {
	return p.name, p.age
}

func main() {
	me := person{
		name: "Antonio",
		age:  19,
	}

	name, age := me.getNameAndAge()

	fmt.Println("My name is", name, "and I'm", age, "years old")
}
