// if - else
package main
import ("fmt")

func main(){
	var a int=3
	var b int=7

	if (b>a){
		fmt.Println("b is greater")
	}else if (a > b){
		fmt.Println("a is greater")
	}
	
	else{
		fmt.Println("default condition executed")
	}
}

// 