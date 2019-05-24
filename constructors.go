package vector

//Init : returns a copy of a initialized vector.
// _Notes:
//		you should set the vecotr equal to the result of it initializing since it returns a copy
//  	EX. "v = v.Init()" is correct usage of this
// _Return:
//		a initialized vector.
func (v *Vector) Init() Vector {
	copyVect := Vector{data: make([]T, 0), size: 0}
	return copyVect
}

//InitWithSlice : returns a copy of a initialized vector that turns the slice into a vector
//_Parameters:
//		elems: T "The slice you want to convert into a vector"
// _Return:
//		a initialized vector.
func (v *Vector) InitWithSlice(elems []T) Vector {
	//it will pull a syntax error if tried to directly initialize through Init()
	var vectClone = Vector{data: make([]T, 0), size: 0}

	for index := 0; index < len(elems); index++ {
		vectClone.PushBack(elems[index])
	}
	return vectClone
}
