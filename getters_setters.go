package vector

/*
File Prolouge:
	This is a file for getters and setters of member variables
	_Format : You must have the getter and the setter. The functions must be aligned
				 in alphabetical order of getter. There are not getEtc() allowed, only
				 Capital Etc() for getters.

				 Setters must have Set then the same name as the getter afterward
				 Ex: At() "getter", SetAt() "setter"
*/

//At : Returns the element of the Vector at a specified index
//_Parameters :
//		index : uint64 "the index of the value you want"
//_Return:
//		the element of the vector at a specified index
func (v *Vector) At(index int) T {
	return v.data[index]
}

//SetAt : Set the element of the Vector at a specified index
//_Parameters :
//		index : uint64 "the index of the value you want"
//		value : int "The element you want to set the value at a specified index instead"
func (v *Vector) SetAt(index int, value T) {
	v.data[index] = value
}

//Size : Returns the size of the vector
//_Return:
//		the size of the vector
func (v *Vector) Size() int {
	return v.size
}
