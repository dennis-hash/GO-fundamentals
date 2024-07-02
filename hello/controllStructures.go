package main

import "fmt"

func controll(){
 printMe("Tumetenga bilioni!")
 var result int = add(1,3)
 fmt.Println(result)
 var i, j int = divide(4,2)
 fmt.Printf("result %v,  remainder %v",i,j)
}

func printMe(printValue string){
 fmt.Println(printValue)
}

func add(int1 int, int2 int) int{
	return int1 + int2
}

func divide(int1 int, int2 int) (int,int){

	if 1==1 && 3==3 {
		fmt.Println("yeah")
	}
	var result int = int1 / int2
	var remainder int = int1 % int2
	return result,remainder
}