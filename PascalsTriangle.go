package main

import "fmt"

func main() {

number:=5
text:='*'

	for i := 0; i < number; i++ {
		for j := 0; j <= i; j++{
			fmt.Print(text)
		}
		fmt.Println()
	}

	
}

//func print(item string){
//	fmt.Print(item string)
// TODO build an easier print method
