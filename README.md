# GoVector

**Version 1.0.0**

go_vector is a container data structure based on C++ std::Vector. go_vector brings functional programming and removes boilerplate code.

#Code samples
*Snippets from the examples/main.go*

Construct an empty vector

```go
var players GoVector.Vector
players = players.Init()
```

Append an element 
```go
players.PushBack(newPlayer)
```

Remove an element
```go
players.RemoveAt(index)
```

Get element at index
```go
players.At(index)
```
Set element at index
```go
players.SetAt(index)
```

Map function
```go
playerCopy := players.FpMap(nil) //putting nil invokes default behavior of just returning a copy of the vector

playerTrophiesDoubled := players.FpMap(func( player GoVector.T ) GoVector.T{
    return player.(Player).trophies * 2
})
```

Reduce function
```go
//noobTrophies is a vector of ints
proTrophieSum := noobTrophies.FpReduce(func(v1 GoVector.T, v2 GoVector.T) GoVector.T {
    return v1.(int) + v2.(int)
}) 
//However, since golang supports the + operator of int and int we could just invoke the default behavior
proTrophieSum := noobTrophies.FpReduce(nil)
```

Filter function
```go
noNoobs := players.FpFilter(func(player GoVector.T) bool {
    //if the player has more than 1000 trophies, add them
    return player.(Player).trophies > 1000
}) //There is no default behavior for filter
```

Sort a vector
```go
//sort the noob trophies from greatest to least, reverse for least to greatest
noobTrophies.SortStruct(func(noob1 GoVector.T, noob2 GoVector.T) bool {
    return noob1.(int) < noob2.(int)
})
```

#Being maintained by
Kode Creer <kodeopensource@gmail.com>

#Contributing
Thanks for being willing to contribute to GoVector. You can join the mailing list for discussion about features. 

#License & Copyright
MIT License
© Kode Creer 2019
