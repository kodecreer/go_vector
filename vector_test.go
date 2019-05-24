package vector

//NOTICE: Please do not add any 3rd party libraries.
import (
	"fmt"
	"testing"
)

//testStruct is a non sensical struct. This is just to see if testVector works with struts
type testStruct struct {
	id    string
	price int
}

func TestInitSize(t *testing.T) {
	/*
		This tests the Size, all Constructors, and Pushback (because init with slice calls it alot)
	*/
	//Testing the initialization of a empty vector
	var testVector Vector
	testVector = testVector.Init()
	if testVector.Size() > 0 {
		t.Error("Init is not empty")
	}

	//Testing converting the slice into a vector

	var testSlice = make([]T, 5)
	for index := 0; index < 5; index++ {
		testSlice[index] = index * index
	}

	testVector = testVector.InitWithSlice(testSlice)

	if testVector.Size() != 5 {
		fmt.Println(testVector)
		t.Errorf("Init with Slice has failed with a size %d vs 5", testVector.Size())
	}

}

func TestPushBack(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()

	//push back an integer
	testVector.PushBack(9)

	if testVector.Size() != 1 && testVector.At(0) != 9 {
		t.Error("Test failed to correctly push back into the vector")
	}

	//Must make new vector to have different types
	var newTestVector Vector
	newTestVector = newTestVector.Init()

	var structure = testStruct{id: "it's over 9000", price: 9001}
	newTestVector.PushBack(structure)

	if newTestVector.At(0).(testStruct).price != 9001 {
		t.Error("Incorrectly adding structs together")
	}
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
func TestSetRemove(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()
	testVector.PushBack(9)
	testVector.PushBack(10)

	testVector.SetAt(0, 9001)
	if testVector.data[0] != 9001 {
		t.Error("Test vector is not setting 9 to be over 9000")
	}

	testVector.RemoveAt(0)
	testVector.PushBack(9)
	if testVector.At(0) != 10 {
		t.Error("Failed to remove value at top")
	}
	testVector.RemoveAt(1)
	if testVector.Size() != 1 {
		t.Error("Failed to remove at end")
	}

}
func TestClear(t *testing.T) {
	var testVector Vector
	testVector = testVector.Init()
	testVector.PushBack(9)
	testVector.PushBack(10)

	testVector.Clear()
	if testVector.Size() > 0 {
		t.Error("Vector failed to clear all elements")
	}
}
