package services

import (
	"fmt"
	"testing"
	"github.com/shwethadia/Testing/src/api/utils/sort"
)

//Integration Test
func TestSortService(t *testing.T){


	elements := sort.GetElements(10)

	fmt.Println("Before Sort",elements)

	Sort(elements)

	//Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}

	fmt.Println("After Sort",elements)

}



func TestSortServiceMoreThan10000(t *testing.T){


	elements := sort.GetElements(20001)

	fmt.Println("Before Sort",elements)

	Sort(elements)

	//Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 20000 {
		t.Error("Last element should be 20000")
	}

	fmt.Println("After Sort",elements)

}

func BenchmarkBubbleSort10k(b *testing.B) {

	//Init
	elements := sort.GetElements(20000)

	//Execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100k(b *testing.B) {

	//Init
	elements := sort.GetElements(100000)

	//Execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}