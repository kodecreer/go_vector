<h1>Getters, setters, append, and remove elements</h1>
Construct an empty vector

```go
var players GoVector.Vector
players = players.Init()
```

Convert a slice into a vector
```go
    players = players.InitWithSlice(playSlice) //slice must be type of GoVector.T
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

Remove last the element of a vector
```go
players.PopBack()
```