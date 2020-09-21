package main

import "fmt"

func binarySearch(arr []int, elem int) {
	low := 0
	high := len(arr) - 1
	var mid int
	flag := false

	for low <= high {
		mid = (high + low) / 2

		if elem < arr[mid] {
			high = mid - 1
		} else if arr[mid] < elem {
			low = mid + 1
		} else {
			flag = true
			break
		}
	}

	if flag == true {
		fmt.Println(elem, "found at", (mid + 1), "position")
	} else {
		fmt.Println(elem, "was not found")
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var elem int
	fmt.Println("Type the number to be searched")
	fmt.Scanf("%d", &elem)

	binarySearch(arr, elem)
}
