// for loop is the only loop available in the go language

package main
import ("fmt")

func main(){
	//var declaration not allowed in for initializer
	for i :=0; i<5; i++ {
		fmt.Println(i)
	}
}


/*
The Range Keyword
The range keyword is used to more easily iterate through the elements 
of an array, slice or map. It returns both the index and the value.

The range keyword is used like this:

for index, value := range array|slice|map {
   // code to be executed for each iteration
}

*/
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}

/*
Output:
0      apple
1      orange
2      banana

*/
