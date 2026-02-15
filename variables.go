package main
import ("fmt")

var a int =2
var b string = "Hello"
var c float32 = 19.29

var e,f,g,h,i int = 4,8,12,16,20 //Mutli-variables decalaration

func main(){
d:=8

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e,f,g,h,i)
}



/* 
** Types of variables in Go:

1. int- stores integers (whole numbers), such as 123 or -123
2. float32- stores floating point numbers, with decimals, such as 19.99 
or -19.99
3. string - stores text, such as "Hello World". String values are 
surrounded by double quotes
4. bool- stores values with two states: true or false

** Declaring Variable type in Go

1. With the var keyword:

Use the var keyword, followed by variable name and type:

var variablename type = value <like> var a int=2

Note: You always have to specify either type or value (or both).

2. With the := sign:

Use the := sign, followed by the variable value

Note: In this case, the type of the variable is inferred from the value 
(means that the compiler decides the type of the variable, based on the 
value).

Note: It is not possible to declare a variable using :=, without 
assigning a value to it.

** Difference Between var and :=

There are some small differences between the var & :=

* var val type:

Can be used inside and outside of functions. 
Variable declaration and value assignment can be done separately	

* :=
Can only be used inside functions
Variable declaration and value assignment cannot be done separately 
(must be done in the same line)

*/