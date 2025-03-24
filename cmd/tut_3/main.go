package main

import "fmt"

func main() {
	intArr:= [3]int32{1, 2, 3}
	fmt.Printf("intArr: %v\n", intArr)

	//slices

	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("intSlice: %v with capactiy %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Printf("intSlice: %v with capacity %v \n", len(intSlice), cap(intSlice))
	fmt.Printf("intArr: %v\n", intArr)

	var intSlice2 []int32 = []int32{2, 3}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("intSlice: %v with capacity %v \n", len(intSlice), cap(intSlice))

	var myMap map[string]uint8 = make(map [string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Eve": 21}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])

	var age,ok = myMap2["Adam"]
	fmt.Println(age, ok)

	age, ok = myMap2["John"]
	fmt.Println(age, ok)

	//for loops

	for name ,age:= range myMap2{
		fmt.Printf("Name: %v\n Age:%v \n", name,age)
	}
}