package main

import (
	"fmt"
)

func main() {

// Addition
	// // With vars declaration
	// var a int = 54
	// var b int = 78
	// var sum int = a+b

	// // Without data type defination
	// a := 45
	// b := 89
	// sum := a+b
	// fmt.Println(sum)

// Comparision
	// a := 97
	// if a >= 50 {
	// 	fmt.Println(a,"Greater then 50.")
	// } else if a >= 40 {
	// 	fmt.Println(a,"Less then 50 but greater then 40.")
	// } else {
	// 	fmt.Println(a,"Less the 40.")
	// }

// Arrays/List
	// // Predefined length
	// data := [2]int{5,4}
	// data[1] = 3
	// fmt.Println(data)

	// // Without length / Slice event
	// data := []int{5,4}
	// data = append(data,3)
	// fmt.Println(data)

// Dictionary / Map
	// data := make(map[int]string)
	// data[1] = "Devendra"
	// data[2] = "Sonali"

	// fmt.Println(data)


// Loops

	// // basic type
	// iterator := 5
	// for iterator>=0 {
	// 	fmt.Println(iterator)
	// 	iterator--
	// }

	// // classic for loop
	// for i:=0;i<5;i++ {
	// 	fmt.Println(i)
	// }

	// // for without any any condition.
	// for {
	// 	fmt.Println("In loop")
	// 	break
	// }

// Switch case
	// // basic
	// i := 1
	// fmt.Println("You have selected",i)
	// switch i {
	// case 1:
	// 	fmt.Println("all")
	// case 2:
	// 	fmt.Println("you")
	// case 3:
	// 	fmt.Println("me")
	// case 4:
	// 	fmt.Println("nobody")
	// }

	// using same case for multiple scenerios
	i := 4

	switch i {
		case 1,3:
			fmt.Println("Odd no")
		case 2,4:
			fmt.Println("Even no")

	}

}



















