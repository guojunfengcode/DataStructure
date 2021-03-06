package main

import (
	"fmt"
)

func QuickSort(values []int) []int{
	len := len(values)
	arr := make([]int, 0, 0)
	if len <= 1 {
		return values
	} else {
		save := values[0]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, save)

		for i := 1; i < len; i++ {
			if values[i] < save {
				low = append(low, values[i])	
			} else if values[i] > save {
				high = append(high, values[i])
			} else {
				mid = append(mid, values[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		arr = append(append(low, mid...), high...)
	}
	copy(values, arr)
	return values	
}

func quickSort_o(values []int, start, end int) {
	if start < end {
		left := start
		right := end

		m := (start + end) / 2
		if values[start] > values[end] {
			values[start], values[end] = values[end], values[start]
		}
		if values[m] > values[end] {
			values[m], values[end] = values[end], values[m]
		}
		if values[m] > values[start] {
			values[m], values[start] = values[start], values[m]
		}
		temp := values[start]
		for left < right {
			for left < right && values[right] >= temp {
				right--
			}
			for left < right && values[left] <= temp {
				left++
			}
			values[left], values[right] = values[right], values[left]
		}
		values[start], values[left] = values[left], values[start]
		quickSort(values, start, left-1)
		quickSort(values, left+1, end)
	}
}

func BinarySearch(arr []int, data int) int {
	left := 0
	right := len(arr) - 1
	if arr[left] == data {
		return left
	}
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	values := []int{4,2,6,1,8,1,4,3,5,2}
	fmt.Println(values)
	fmt.Println("=======QuickSort=========")
	QuickSort(values)
	fmt.Println(values)
	fmt.Println("=================")
	fmt.Println(BinarySearch(values, 6))
}
