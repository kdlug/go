package main

import "fmt"

func main() {

	games := []string{"pokemon", "sims"}
	newGames := []string{"pacman", "doom", "heroes"}

	newGames = games
	fmt.Println(games)
	fmt.Println(newGames)

	games = nil
	fmt.Println(games)
	fmt.Println(newGames)

}
