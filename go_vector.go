package GoVector

//T is meant to signify that an interface is supposed to act like a psuedo generic.
// Mainly mean to undo confusion in the codebase
type T interface{}

//AltT is the same as T, but it is used so funcitons can return alternating types
type AltT interface{}

//Vector is a list based on C++ Vector with functional programming functions
type Vector struct {
	data []T
	size int
}

//PushBack : Adds an item to the end of the vector, just like C++
// _Parameters:
//			element: T "the value you want to add"
func (v *Vector) PushBack(element T) {
	v.data = append(v.data, element)
	v.size++
}

//RemoveAt : removes an item from the vector at a specified index
// _Parameters:
//			index: int "the index of the value you want to remove it from"
func (v *Vector) RemoveAt(index int) {
	v.data = append(v.data[:index], v.data[index+1:]...)
	v.size--
}

//Clear : removes all elements from the vector
func (v *Vector) Clear() {
	for v.Size() > 0 {
		v.RemoveAt(0)
	}
}

//SortStruct : sorts the vecotr
//_Notes: Use this to sort your structs. You must implment a compare funciton that runs like this
// 			func(v1, v2){ return v1.compareValue < v2,compareValue }
// 			You can define an inline function that returns the opposite of your sort in you Struct if you want it to be reversed
// _Parameters:
//			sortOperator: T "the reference to the function that would be used as a comparison operator"
func (v *Vector) SortStruct(function func(v1 T, v2 T) bool) {
	//Golang is picky about uint64 and ints
	for i := 0; i < int(v.size-1); i++ {
		for j := i + 1; j < int(v.size); j++ {
			if function(v.data[i], v.data[j]) && i != j {
				temp := v.data[j]
				v.data[j] = v.data[i]
				v.data[i] = temp
			}
		}
	}
}
