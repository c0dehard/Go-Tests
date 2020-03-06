package main
import(
	"bufio"
	"os"
)

func main(){

  /* IMPORTANT
  
   :=  Declaration And Assignment
    =  Assignment Only
    
   */
   
	f,e := os.Open("yourFile.extension")
	s := bufio.NewScanner(f)
	if e != nil{
		print("Error")
		return
	}
	
	for s.Scan(){
		println(s.Text())
	}
}
