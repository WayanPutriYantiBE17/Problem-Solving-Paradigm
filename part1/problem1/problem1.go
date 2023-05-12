package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) interface{} {
    var x, y, z int
    for x = 0; x < 1000; x++ {
        for y = 0; y < 1000; y++ {
            for z = 0; z < 1000; z++ {
                if x+y+z == a && b == x*y*z && c == int(math.Pow(2, float64(x))+math.Pow(2, float64(y))+math.Pow(2, float64(z))) {
                    return []int{x, y, z}
                }
            }
        }
    }
    return "No solution"
}
func main() {
	fmt.Println(SimpleEquations(1, 2, 3))  //no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // [1,2,3]
}
