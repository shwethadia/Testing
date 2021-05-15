package services

import(

	"github.com/shwethadia/Testing/src/api/utils/sort"
)
func Sort(elements []int){

//	sort.BubbleSort(elements)
if len(elements) <= 20000{

	sort.BubbleSort(elements)
	return 
}
sort.Sort(elements)
}