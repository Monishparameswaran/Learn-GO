# Slice 
## what is a slice in go?
   The slice type is an abstraction built on top of Go’s array type, and so to understand slices we must first understand arrays.slice is an abstraction of array
and efficient software systems.

    slice are more flexible and powerful .variable length
    slice are also index-based and have a asixe but is resized when needed
    slice are much flexible then arrays
    
## Intialise the slice
>   letters := [ ]string{"M", "o", "n", "i","s","h"}.

### How to add elements to the slice
>  letters = append(letters,"name: Monish")

### len and cap in slices
 
 
 A slice can be created with the built-in function called make, which has the signature,


 
``` 
func make([ ]T, len, cap) [ ]T{

 }
```

``` 
 name:=make([ ]string,0,10) 
```

 len(letters)   //gives the length 


### what is length 
  It defines how many elements are there in slice currently
 
### What is capacity
 Capacity represents the maximum number of elements that the slice can hold without reallocating the underlying array.
  
   frequent reallocation can affect perfomance
