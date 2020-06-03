package main

import "./constructor"

func main() {
	p := constructor.Player("C")
	for i := 0; i < 4; i++ {
		println(p.Item[i].Name)
	}
}
