package playGround

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"unsafe"
)

type Pointer struct {
	Name string
}

func TestPointer(t *testing.T) {
	var target Pointer
	fmt.Println("size of ", unsafe.Sizeof(target))
	target = Pointer{
		Name: "aaa",
	}

	fmt.Printf("address , %p \n", &target)
	getAnyStruct(&target)
	fmt.Printf("address , %p \n", &target)
	fmt.Println("pointer ", target)
}

func TestCopyStructure(t *testing.T) {
	target := &Pointer{
		Name: "aaa",
	}

	tt := *target
	tt.Name = "bbb"
	fmt.Println("target :", target)
	fmt.Println("tt :", tt)

}

func getAnyStruct(i interface{}) {
	converted := i.(*Pointer)
	converted.Name = "bb"
}

func TestUnmarshal(t *testing.T) {
	data := []byte("{\"test\": \"test\", \"TEST\": \"TEST\"}")
	var d struct {
		LowerCase string `json:"test"`
		UpperCase string `json:"TEST"`
	}
	err := json.Unmarshal(data, &d)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(d.LowerCase)
	fmt.Println(d.UpperCase)

	// Output:
	// test
}
