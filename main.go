package main

func main() {
	//Create objects
	game1 := game{
		title: "Minecraft",
		price: 5,
	}
	game2 := game{
		title: "World of Warcraft",
		price: 19,
	}
	game3 := game{
		title: "Elite Dangerous",
		price: 54,
	}

	var items []*game
	items = append(items, &game1, &game2, &game3)

	my := list(items)
	my.print()
}
