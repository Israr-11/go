//Example 1
package main
import ("fmt")

func main(){
var a = map[string]string{"car":"ford", "model":"2025"}

fmt.Printf("The value %v", a)
}

/*
--> Maps are used to store data values in key:value pairs.

--> Each element in a map is a key:value pair.

--> A map is an unordered and changeable collection that does not allow duplicates.

--> The length of a map is the number of its elements. You can find it using 
the len() function.

--->The default value of a map is nil.Maps hold references to an underlying hash 
table. Go has multiple ways for creating maps.

Syntax:

var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

*/

// Example 2

//Making the map using the make keyword

// var a = make(map[KeyType]ValueType)
// b := make(map[KeyType]ValueType)

package main
import ("fmt")

func main() {
  var a = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"
                                 // a is no longer empty
  b := make(map[string]int)
  b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

//Example 3

// Access Map Elements

package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}