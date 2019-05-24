package main

import (
	"fmt"

	GoVector "github.com/kodecreer/go_vector"
)

type Player struct {
	id       int
	name     string
	trophies int
}

func main() {
	var players GoVector.Vector
	playSlice := make([]GoVector.T, 20)
	for i := 0; i < 20; i++ {
		playerName := fmt.Sprintf("NoobSlayer %d", i)
		player := Player{id: i, name: playerName, trophies: i * 100}
		playSlice[i] = player
	}
	//Or, if I want to initialize an empty  vector, I call Init() instead
	players = players.InitWithSlice(playSlice)

	//Oh no, I messed up with the first index
	playerTypo := Player{id: 0, name: "King Gopher over 9000", trophies: 9001}
	//nothing like a easy fix
	players.SetAt(0, playerTypo)

	//putting nil in the parameter invokes the default behavior for the map funciton
	//which just returns copy of the vector
	playerCopy := players.FpMap(nil)
	fmt.Printf("Copy size is %d vs the original size of %d", playerCopy.Size(), players.Size())

	//Declaring an inline function that act as a closure
	noNoobs := players.FpFilter(func(player GoVector.T) bool {
		fmt.Println(player.(Player).name)
		return player.(Player).trophies > 1000
	})
	noobTrophies := noNoobs.FpMap(func(player GoVector.T) GoVector.T {
		//If I wanted to return a whole copy of an array, then I just return player and set it equal to map function
		return player.(Player).trophies
	})
	proTrophieSum := noobTrophies.FpReduce(func(v1 GoVector.T, v2 GoVector.T) GoVector.T {
		return v1.(int) + v2.(int)
	})

	fmt.Println(proTrophieSum)
	//Let's empty the vecotr
	players.Clear()

}
