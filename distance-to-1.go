/*

Problem statement:

On a positive integer, you can perform any one of the following 3 steps. 1.) Subtract 1 from it. ( n = n - 1 )  , 2.) If its divisible by 2, divide by 2. ( if n % 2 == 0 , then n = n / 2  )  , 3.) If its divisible by 3, divide by 3. ( if n % 3 == 0 , then n = n / 3  ). Now the question is, given a positive integer n, find the minimum number of steps that takes n to 1

eg: 1.)For n = 1 , output: 0       2.) For n = 4 , output: 2  ( 4  /2 = 2  /2 = 1 )    3.)  For n = 7 , output: 3  (  7  -1 = 6   /3 = 2   /2 = 1 )

*/

package main

import (
	"fmt"
)

const MAX_ITR = 100
const INFINITY = 1<<31 - 1
const NUM = 9

func min(x int, y int, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}

func calcMin(n int, minMap *[]int) int {
	var x int = INFINITY
	var y int = INFINITY
	var z int = INFINITY
	if n%3 == 0 {
		x = n / 3
		x = (*minMap)[x] + 1
	}
	if n%2 == 0 {
		y = n / 2
		y = (*minMap)[y] + 1
	}
	z = (*minMap)[n-1] + 1
	return min(x, y, z)
}

func main() {
	var minMap []int
	minMap = make([]int, MAX_ITR)
	minMap[0] = -1
	minMap[1] = 0
	minMap[2] = 1
	minMap[3] = 1
	for i := 4; i < MAX_ITR; i++ {
		minMap[i] = calcMin(i, &minMap)
	}
	fmt.Println(INFINITY)
	fmt.Println("Min distance from", NUM, "to 1:", minMap[NUM])
}
