package main

import "fmt"

type Pessoa struct {
	name string
}

func (p *Pessoa) falar() {
	fmt.Println("Me chamo", p.name)
}

type Humano interface {
	falar()
}

func dizerAlgo(h Humano) {
	h.falar()
}

func main() {
	eu := Pessoa{"Antonio"}
	dizerAlgo(&eu)
}
