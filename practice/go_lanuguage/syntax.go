package main
import ("fmt")

func main(){
	fmt.Println("Hello")
}

/*
Line 1: In Go, every program is part of a package. We define this using the package keyword.
In this example, the program belongs to the main package.

Line 2: import ("fmt"), lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more 
readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. 
It is used to output/print text. In our example it will output "Hello World!".
*/