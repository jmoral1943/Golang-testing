package services

import "github.com/jmoral1943/udemy/go_testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) > 10000 {
		sort.Sort(elements)
		return
	}
	sort.BubbleSort(elements)
}