/*

Go has three functions to output text, Print(), Println(), Printf()

1. Print()

The Print() function prints its arguments with their default format. Its
mean that no formatting and nothing is done. Printed as it is, without
whitespace and stuff.

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}

Output: HelloWorld


2. Println()

The Println() function is similar to Print() with the difference that a 
whitespace is added between the arguments, and a newline is added at the 
end:

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}

Output: Hello World

3. Printf()

The Printf() function first formats its argument based on the given 
formatting verb and then prints them.

Here we will use two formatting verbs:

A. %v is used to print the value of the arguments
B. %T is used to print the type of the arguments

package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}

Output:

i has value: Hello and type: string
j has value: 15 and type: int

********************************************************************
****************Common output formatting****************************
********************************************************************
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign
********************************************************************
*/