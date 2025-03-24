package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 30, gallons: 10, ownerInfo: owner{name: "John Doe"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
}