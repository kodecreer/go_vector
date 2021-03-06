# GoVector

**Version 1.2.0**

go_vector is a container data structure based on C++ std::Vector. go_vector brings functional programming and removes boilerplate code.

<h1>Code samples</h1>
*Snippets from the documentation/examples/main.go*

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

<h1>Being maintained by</h1>
Kode Creer <a href="mailto:kodeopensource@gmail.com">Contact / mailing list email</a>

<h1>Contributing</h1>
Thanks for being willing to contribute to GoVector. You can join the mailing list for discussion about features . There is the dev branch, which is where you would modify and make changes to the librairy. Each feature must have a test to ensure it works with that follows the documentation convention. You can also help out the community by adding more example projects using GoVector that act like tutorials or film them on Youtube. 

To save you a stack overflow search here is how you push the dev branch to the release
```
git push origin dev:master
```
<h1>License & Copyright</h1>
MIT License
© Kode Creer 2019
