package govector

/* File Prolouge:
Because there is no need to access private variables here, then we use Unit
*/
import (
	"testing"
)

//testStruct is a non sensical struct. This is just to see if testVector works with struts
type testStruct struct {
	id    string
	price int
}

func TestMap(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()

	for i := 0; i < 100000; i++ {
		testVector.PushBack(i)
	}

	if testVector.Size() != 100000 {
		t.Errorf("Not enough Elements : %d", testVector.Size())
	}
	//testing if  the default behavior works
	defaultCopy := testVector.FpMap(nil)

	//testing if the closured like behavior works
	closureCopy := testVector.FpMap(func(num T) T {
		return num.(int) * 2
	})

	for i := 0; i < closureCopy.Size(); i++ {
		if closureCopy.At(i) != testVector.At(i).(int)*2 {
			t.Error("Mismatching values with closure copy")
		}
		if defaultCopy.At(i) != testVector.At(i) {
			t.Error("Mismatching values with default copy")
		}
	}

}

func TestFpReduce(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()

	for i := 0; i < 5; i++ {
		testVector.PushBack(i)
	}
	if testVector.Size() != 5 {
		t.Errorf("Not enough Elements : %d", testVector.Size())
	}
	//testing default behavior
	var reduced = testVector.FpReduce(nil).(int)

	if reduced != 10 {
		t.Errorf("Reduced supposed to be 10 not %d", reduced)
	}

	//testing struct behavior
	var structVector Vector
	structVector = structVector.Init()

	var structure1 = testStruct{id: "it's half of over 9000", price: 4501}
	var structure2 = testStruct{id: "it's half of 9000", price: 4500}

	structVector.PushBack(structure1)
	structVector.PushBack(structure2)

	var structReduced = structVector.FpReduce(func(v1 T, v2 T) T {
		return v1.(testStruct).price + v2.(testStruct).price
	}).(int)

	if structReduced <= 9000 {
		t.Error("Struct Reduce is not over 9000!!!!! >:(")
	}
}

func TestFpFilter(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()

	//size should be 5
	for i := 0; i < 5; i++ {
		testVector.PushBack(i)
	}

	var filteredVector = testVector.FpFilter(func(val T) bool {
		return val.(int) < 2
	})
	if filteredVector.Size() != 2 {
		t.Errorf("Vecotr size supoosed to be 2 not %d", filteredVector.Size())
	}
	if filteredVector.At(0) != 0 && filteredVector.At(1) != 1 {
		t.Error("Vector incorrectly filtered itself")
	}
	//testing struct behavior
	var structVector Vector
	structVector = structVector.Init()

	var structure1 = testStruct{id: "it's half of over 9000", price: 4501}
	var structure2 = testStruct{id: "it's half of 9000", price: 4500}

	structVector.PushBack(structure1)
	structVector.PushBack(structure2)

	filteredVector = structVector.FpFilter(func(val T) bool {
		return val.(testStruct).price > 4500
	})

	if filteredVector.Size() != 1 {
		t.Error("Filtered Vector incorrectly filtered vector")
	}

}

func TestPopBack(t *testing.T) {
	/*
		This tests the popback. The result shoudl be left with a vector with a size of 1 and
	*/
	var testVector Vector
	testVector = testVector.Init()

	testVector.PushBack(9001)
	testVector.PushBack(80)

	testVector.PopBack()
	if testVector.At(0) != 9001 {
		t.Error("Value removed is incorrect ")
	}
	if testVector.Size() > 1 {
		t.Error("testVector did not remove element")
	}
}
func TestFpIndexOf(t *testing.T) {
	/*
		This sees whether it returns the correct index in the functional programing indexOf funciton
	*/
	var testVector Vector
	testVector = testVector.Init()

	testVector.PushBack(9001)
	testVector.PushBack(80)
	testVector.PushBack(4500)
	testVector.PushBack(324)

	//called index1 because it's the desired index
	index1 := testVector.FpIndexOf(func(value T) bool {
		return value.(int) == 80
	})

	if index1 != 1 {
		t.Error("Wrong index returned")
	}
}
