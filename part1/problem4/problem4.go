package main

import "fmt"

func BinarySearch(array []int, x int) int {
	// your code here
var kiri = 0
var kanan = len(array)-1


for kiri <= kanan{
	var tengah = (kanan + kiri)/2	
	if x > array[tengah]{
		kiri = tengah+1
	}else if x< array[tengah]{
		kanan = tengah-1
	}else if x == array[tengah]{
		return tengah
	}

	}
return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))                  //2
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))                 //3
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))  //6
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)) //-1
}
