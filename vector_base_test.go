package govector

/*File Prolouge:
This is where all the base vector functionality are tested
*/
import (
	"fmt"
	"testing"
)

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
	/*
		Tests the pushback
	*/
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
func TestSetRemove(t *testing.T) {

	/*
		This tests SetAt and RemoveAt
	*/
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
	/*
		This tests Clear
	*/
	var testVector Vector
	testVector = testVector.Init()
	testVector.PushBack(9)
	testVector.PushBack(10)

	testVector.Clear()
	if testVector.Size() > 0 {
		t.Error("Vector failed to clear all elements")
	}
}
