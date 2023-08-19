package main

import (
	"fmt"
	"github.com/alhaqz/puppy"
)

func main() {
	s1, s2, s3, s4 := puppy.Bark(), puppy.Barks(), puppy.BigBark(), puppy.BigBarks()

	fmt.Println(s1, s2, s3, s4)
	
} 