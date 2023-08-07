//Given two sorted lists, write a function to merge them into a single sorted list.

package main

import "fmt"

func mergeSortedList(list1, list2 []int) []int {
	mergedList := make([]int, 0)

	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			mergedList = append(mergedList, list1[i])
			i++
		} else {
			mergedList = append(mergedList, list2[j])
			j++
		}
	}
	for i < len(list1) {
		mergedList = append(mergedList, list1[i])
		i++
	}

	for j < len(list2) {
		mergedList = append(mergedList, list2[j])
		j++
	}
	return mergedList
}

func main() {
	sortedList1 := []int{1, 3, 5, 7, 9}
	sortedList2 := []int{2, 4, 6, 8, 10}

	mergedlist := mergeSortedList(sortedList1, sortedList2)

	fmt.Println("Sorted List1: ", sortedList1)
	fmt.Println("Sorted List2: ", sortedList2)
	fmt.Println("Sorted Merged List is: ", mergedlist)
}
