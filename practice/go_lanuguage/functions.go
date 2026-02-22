package main
import ("fmt")

//The fname in this case is called as parameter
func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}


//Return Type

package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

//Recursion

//In the following example, testcount() is a function that calls itself. 
// We use the x variable as the data, which increments with 1 (x + 1) 
// every time we recurse. The recursion ends when the x variable equals 
// to 11 (x == 11). 


package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}