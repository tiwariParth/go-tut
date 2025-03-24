package main

import "fmt"

func main(){
	var p *int32
	var i int32

	p = new(int32)
	p = &i
	*p = 23
	fmt.Printf("p: %v\n", *p)
	fmt.Printf("i: %v\n", i)
}

