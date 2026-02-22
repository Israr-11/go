
/*
Unlike arrays, the length of a slice can grow and shrink as 
you see fit.

slice_name := []datatype{values}

myslice := []int{1,2,3}
The code above declares a slice of integers of length 3 and also the 
capacity of 3.

***** In Go, there are two functions that can be used to return the 
length and capacity of a slice:

---> len() function - returns the length of the slice (the number of 
elements in the slice)
----> cap() function - returns the capacity of the slice (the number of 
elements the slice can grow or shrink to)

**** Slice Function

The make() function can also be used to create a slice.

Syntax

slice_name := make([]type, length, capacity)


************* Memory Efficiency ***********
 
When using slices, Go loads all the underlying elements into the memory.

If the array is large and you need only a few elements, it is better to 
copy those elements using the copy() function.

The copy() function creates a new underlying array with only the required 
elements for the slice. This will reduce the memory used for the program. 

**** Syntax
copy(dest, src)

The copy() function takes in two slices dest and src, and copies data 
from src to dest. It returns the number of elements copied.*******




*/