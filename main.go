package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group I's Week 4 Project!")
}

// Rohit Sthapit
// BubbleSort function sorts a slice of integers using the Bubble Sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
