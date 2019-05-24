package vector

//FpMap : Returns a copy of the vector that can be altered or not through the inline function
//	_Parameters:
//		function: T "A funciton with one arugument that matches the datatype of the vector's elements,
// it must return a value that is the same datatype as the slice values or it will crash the program"
//	_Returns:
//		T "The filtered vector that is the same datatype as the vector"
//
func (v *Vector) FpMap(function func(T) T) Vector {
	var dataCopy Vector

	dataCopy = dataCopy.Init()

	for index := 0; index < len(v.data); index++ {
		if function != nil {
			mapFunc := function(v.data[index]) //v.Mapable(v.data[index])
			dataCopy.PushBack(mapFunc)
		} else {
			dataCopy.PushBack(v.data[index])
		}
	}
	return dataCopy
}

//FpReduce : reduces a whole value into one and treat it as a single sum from the vector
//	_Parameters:
//		function: T "A funciton with one arugument that matches the datatype of the vector's elements,
// it must return a value that is the same datatype as the slice values or it will crash the program
//	_Returns:
//		T "The accumulate value that is the same datatype as the vector"
//
func (v *Vector) FpReduce(function func(T, T) T) T {
	sum := v.data[0]
	for index := 1; index < len(v.data); index++ {
		if function != nil {
			sum = function(sum, v.data[index])
		} else {
			switch sum.(type) {
			//For those of you wondering why there is a sum temp, it is because
			//Go doesn't allow use to += with sum.(int), which is something we are trying to do
			//THis seems like a little hack, but at least not a painful one....
			case int:
				sumTemp := sum.(int)
				sumTemp += v.data[index].(int)
				sum = sumTemp
				break
			case int16:
				sumTemp := sum.(int16)
				sumTemp += v.data[index].(int16)
				sum = sumTemp
			//The documentation said the int32 is not an alias for int
			case int32:
				sumTemp := sum.(int32)
				sumTemp += v.data[index].(int32)
				sum = sumTemp
			case int64:
				sumTemp := sum.(int64)
				sumTemp += v.data[index].(int64)
				sum = sumTemp
			case float32:
				sumTemp := sum.(float32)
				sumTemp += v.data[index].(float32)
				sum = sumTemp
			case float64:
				sumTemp := sum.(float64)
				sumTemp += v.data[index].(float64)
				sum = sumTemp
			case string:
				sumTemp := sum.(string)
				sumTemp += v.data[index].(string)
				sum = sumTemp
			default:
				panic(`Nil function failed to accumulate vector values, possible reasons \n
				* Vector has mismatching types \n
				* Attempting to add Types that don't have the + operator supported by default in Golang`)
			}
		}
	}
	return sum
}

//FpFilter : Returns a filtered copy of the vector that satifys the given conditions
//	_Parameters:
//		function: T "A funciton with one arugument that matches the datatype of the vector's elements,
// it must return a boolean or it will crash the program"
//	_Returns:
//		T "The filtered vector that is the same datatype as the vector"
//
func (v *Vector) FpFilter(function func(T) bool) Vector {
	var dataCopy Vector
	dataCopy = dataCopy.Init()
	for index := 0; index < len(v.data); index++ {
		if function(v.data[index]) {
			dataCopy.PushBack(v.data[index])
		}
	}
	return dataCopy
}
