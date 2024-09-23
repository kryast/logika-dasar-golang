package main

import "fmt"

func main() {

	/**
	buatlah output
	*
	**
	***
	****
	*****

	dari input 5

	*/

	input := 5

	for i := 1; i <= input; i++ {
		for j := 1; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println("*")
	}

	/**
	buatlah output
	    *
	   ***
	  *****
	 *******
	*********

	dari input 5

	*/

	n := 5

	for i := 1; i <= n; i++ {
		for a := 1; a <= n-i; a++ {
			fmt.Print(" ")
		}
		for j := 1; j < i+(i-1); j++ {
			fmt.Print("*")
		}
		fmt.Println("*")
	}
}
