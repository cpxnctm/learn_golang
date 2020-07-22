package main

import "fmt"

func main(){

fmt.Println(foo())

}


func foo()int{
	x :=[]int{25,32}
	y := 0
	for _,v := range x{
		y = v + v
	}
	return y
}