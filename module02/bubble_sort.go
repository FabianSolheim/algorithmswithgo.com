package module02

import (
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list


func BubbleSortInt(list []int)  {
	for i := 0; i < len(list); i++ {
		
		for j := 0; j < len(list) - i - 1; j++ {
			if(list[j] > list[j + 1]) {
				temp := list[j]
				list[j] = list[j + 1]
				list[j + 1] = temp
			}
		}
	}
}


// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
}
