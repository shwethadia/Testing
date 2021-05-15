package sort

import (
	"fmt"
	"testing"
)

//Unit Test
func TestBubbleSortOrderASCEND(t *testing.T) {
	//∏
	elements := []int{9, 5, 7, 3, 1, 2, 4, 6, 8, 0}

	//Before sort
	fmt.Println("Before sort", elements)

	//Execution
	BubbleSort(elements)

	//Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}

	fmt.Println("After Sort", elements)

}

func TestBubbleSortAlreadySorted(t *testing.T) {

	//Init
	elements := []int{5, 4, 3, 2, 1}

	//Before sort
	fmt.Println("Before sort", elements)

	//Execution
	BubbleSort(elements)
	fmt.Println("After Sort", elements)
}


func TestSortIncreasingOrder(t *testing.T) {

	//Init
	elements := GetElements(10)

	//Before sort
	fmt.Println("Before sort", elements)

	timeoutChan := make(chan bool,1)
	defer close(timeoutChan)

	go func(){

		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func(){
		time.Sleep(500 * time.Millisecond)
		timeoutChan <-true
	}()

	if  <- timeoutChan {

		assert.Fail("Bubble sort took more than 500 ms")
		return 
	}

	

	//After sort
	fmt.Println("After sort", elements)

	//Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}
}


func TestSortIncreasingOrder(t *testing.T) {

	//Init
	elements := GetElements(10)

	//Before sort
	fmt.Println("Before sort", elements)

	Sort(elements)

	//After sort
	fmt.Println("After sort", elements)

	//Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}
}



/* func BenchmarkBubbleSort(b *testing.B) {

	//Init
	elements := []int{9, 5, 7, 3, 1, 2, 4, 6, 8, 0}

	//Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} */

func BenchmarkBubbleSort(b *testing.B) {

	//Init
	elements := GetElements(10000)

	//Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

/* func BenchmarkSort(b *testing.B) {

	//Init
	elements := []int{9, 5, 7, 3, 1, 2, 4, 6, 8, 0}

	//Execution
	for i := 0; i < b.N; i++ {
	
		Sort(elements)
	}
} */

func BenchmarkSort(b *testing.B) {

	//Init
	elements := GetElements(10000)

	//Execution
	for i := 0; i < b.N; i++ {
	
		Sort(elements)
	}
}
//Benchmark allows us to perform these type of comparison between different approaches for the same problem


