// strings, rune, bytes

package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "resume"

	var indexed = myString[0]
	fmt.Printf("%v, %T",indexed, indexed)

	//string building

	var myString1 = []string{"h", "e", "l", "l", "o"}
	var myString2 = []string{"w", "o", "r", "l", "d"}
	var catString = ""
	for i:=range myString1{
		catString += myString1[i]
	}
	for i:= range myString2{
		catString += myString2[i]
	}
	fmt.Println("\n",catString)

	//strings package
	var strBuilder strings.Builder
	for i:= range myString1{
		strBuilder.WriteString(myString1[i])
	}
	for i:= range myString2{
		strBuilder.WriteString(myString2[i])
	}

	catString = strBuilder.String()
	fmt.Println(catString)


}