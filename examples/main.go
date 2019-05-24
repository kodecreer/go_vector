package main

import (
	"fmt"

	vector "github.com/kodecreer/go_vector"
)

type Player struct {
	id       int
	name     string
	trophies int
}

func main() {
	var players vector.Vector
	playSlice := make([]vector.T, 20)
	for i := 0; i < 20; i++ {
		playerName := fmt.Sprintf("NoobSlayer %d", i)
		player := Player{id: i, name: playerName, trophies: i * 100}
		playSlice[i] = player
	}
	//Or, if I want to initialize an empty  vector, I call Init() instead
	players = players.InitWithSlice(playSlice)

	//Declaring an inline function that act as a closure
	players.FpMap(func(player vector.T) vector.T {
		// fmt.Println(player.(Player).name)
		//If I wanted to return a whole copy of an array, then I just return player and set it equal to map function
		return nil
	})

	noNoobs := players.FpFilter(func(player vector.T) bool {
		// fmt.Println(player)
		return player.(Player).trophies > 1000
	})
	noobTrophies := noNoobs.FpMap(func(player vector.T) vector.T {
		//If I wanted to return a whole copy of an array, then I just return player and set it equal to map function
		return player.(Player).trophies
	})
	proTrophieSum := noobTrophies.FpReduce(func(v1 vector.T, v2 vector.T) vector.T {
		return v1.(int) + v2.(int)
	})

	fmt.Println(proTrophieSum)

}
