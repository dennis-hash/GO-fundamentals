package main

import (
	"fmt"
	"unicode/utf8"
)

func dataTypes() {
	var intNum int //int16, int32 int64 uint
	intNum++
	fmt.Println(intNum)

	var floatNum float64 = 12.0
	fmt.Println(floatNum)

	var myString string = "hello world" + "!"
	myNewString := "wow"
	fmt.Println(myString, myNewString)

	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	const myBool bool = true
	fmt.Println(myBool)

}
