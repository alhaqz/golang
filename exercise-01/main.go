package main

import (
	"fmt"
	// "math/rand"
	// "math"
)

// func add(x, y int) int {
// 		return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x,y int) {
// 	x = sum * 12
// 	y = sum - x
// 	return
// }

// var c, python, java bool

const (
    B int = 1 << (iota * 10) // 1 << (0 * 10) = 1
    KB  // 1 << (1 * 10) = 1024
    MB  // 1 << (2 * 10) = 1048576
    GB  // 1 << (3 * 10) = 1073741824
    TB  // 1 << (4 * 10) = 1099511627776
)


func main() {
	fmt.Printf("B: %d \t\t\t %b\n", B, B)
	fmt.Printf("KB: %d \t\t %b\n", KB, KB)
	fmt.Printf("MB: %d \t\t %b\n", MB, MB)
	fmt.Printf("GB: %d \t\t %b\n", GB, GB)
	fmt.Printf("TB: %d \t %b\n", TB, TB)
	
	// fmt.Printf("%d \t %b\n", 1, 1)
	// fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	// fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	// fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	// fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	// fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	// fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
	
	
	// var i int
	// fmt.Println(i, c, python, java )
	// const name, age = "Budi", 23
	// fmt.Println(add(32, 13))
	// fmt.Println(swap("lalala", "yeyeye"))
	// fmt.Println(split(12))
	// fmt.Printf("%s is %d year's old\n", name, age)
	// fmt.Println(rand.Intn(3))
	// fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	// fmt.Printf("Now you have %g problems. \n", math.Pi)
}