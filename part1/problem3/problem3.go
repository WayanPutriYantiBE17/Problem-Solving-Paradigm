package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) string {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	if len(dragonHead) > len(knightHeight){
		return "knight fall"
	}

	n := 0
	for _, D := range dragonHead {
		for i, H := range knightHeight {
			if H >= D {
				n +=H
				knightHeight = append(knightHeight[:i], knightHeight[i+1:]...)
				break
			}
		}
	}
	return fmt.Sprint(n)
}

func main() {
	fmt.Println(DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}))    //11
	fmt.Println(DragonOfLoowater([]int{5, 10}, []int{5}))         // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})) // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})) // 10
}
