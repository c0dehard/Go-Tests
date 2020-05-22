package main

func lifo(){
	if 0 < 1 {
		defer println("O")
	} 
	if "GNU" != "UNIX"{
		defer print("F")
	}
	defer print("I")
	two := 2
	three := 3
	if two != three{
		defer print("L")
	}
	println("Defer is always executed only after the function call")
}

func main(){
	lifo()
}
