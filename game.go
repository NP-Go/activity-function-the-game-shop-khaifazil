package main

import "fmt"

//Declare Struct for Game

type game struct {
	title string
	price float64
}

//Declare method for Game

func (g game) print() {
	fmt.Println(g.title, "$", g.price)
}
