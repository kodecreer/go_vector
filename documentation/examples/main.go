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

	//Construct a slice to later be converted into a vector
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

	//let get the last player from the vector
	lastIndex := players.Size() - 1

	lastPlayer := players.At(lastIndex)

	newPlayer := Player{id: lastPlayer.(Player).id + 1, name: "Gopher minion", trophies: 5045}
	//One player just joined the game, lets add them on
	players.PushBack(newPlayer)

	//putting nil in the parameter invokes the default behavior for the map funciton
	//which just returns copy of the vector
	playerCopy := players.FpMap(nil)
	fmt.Printf("Copy size is %d vs the original size of %d\n", playerCopy.Size(), players.Size())

	//Declaring an inline function that act as a closure
	noNoobs := players.FpFilter(func(player GoVector.T) bool {
		fmt.Println(player.(Player).name)
		return player.(Player).trophies > 1000
	})
	noobTrophies := noNoobs.FpMap(func(player GoVector.T) GoVector.T {
		//If I wanted to return a whole copy of an array with just the player trophies, then I just return player and set it equal to map function
		return player.(Player).trophies
	})
	fmt.Println(noobTrophies)
	//sort the noob trophies from greatest to least, reverse for least to greatest
	noobTrophies.SortStruct(func(noob1 GoVector.T, noob2 GoVector.T) bool {
		return noob1.(int) < noob2.(int)
	})
	fmt.Println(noobTrophies)

	//This is the structor of what we normally would do if the + operator isn't supported by a struct or data type
	// proTrophieSum := noobTrophies.FpReduce(func(v1 GoVector.T, v2 GoVector.T) GoVector.T {
	// 	return v1.(int) + v2.(int)
	// })
	//Since golang supports the + operator of int and int we could just invoke the default behavior
	proTrophieSum := noobTrophies.FpReduce(nil)
	//which is the same as

	fmt.Println(proTrophieSum)
	//Let's empty the vecotr
	players.Clear()

}
