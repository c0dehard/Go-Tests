package main

//import "fmt"

func main() {

number:=5
text:="*"

	for i := 0; i < number; i++ {
		for j := 0; j <= i; j++{
			print(text)
		}
		println()
	}
}

/* fmt.Print isn't the only way!
   I dont need fmt right now
   for this little script*/
