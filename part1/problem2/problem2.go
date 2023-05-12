package main

import (
	"fmt"
	"sort"
)

func MoneyChange(money int) []int {
	// your code here
	M :=[]int{1, 10, 20, 50, 100, 200, 500, 1000, 2000,5000, 10000}
	sort.SliceStable(M, func(i, j int) bool {
		return M[i] > M[j]
	})
	result :=[]int{}
	for i:=0;i<len(M);i++{
		for money>=M[i]{
			result = append(result, M[i])
			money = money-M[i]

		}
	}
	return result
}

func main() {
	fmt.Println(MoneyChange(123))   //[100 20 1 1 1]
	fmt.Println(MoneyChange(432))   //[200 200 20 10 1 1]
	fmt.Println(MoneyChange(543))   //[500 20 20 1 1 1]
	fmt.Println(MoneyChange(7752))  //[5000 2000 500 200 50 1 1]
	fmt.Println(MoneyChange(15321)) //[10000 5000 200 100 20 1]
}
