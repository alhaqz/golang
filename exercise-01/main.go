package main

import (
	"fmt"
	"github.com/alhaqz/golang/puppy"
)

func main() {
	s1, s2 := puppy.Bark(), puppy.Barks()

	fmt.Println(s1, s2)
	
}