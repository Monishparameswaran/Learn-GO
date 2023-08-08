# What is map

A map in Go is a built-in data structure that allows you to store and retrieve values using keys. It's similar to a dictionary or hash table in other programming languages. In a map, each key must be unique, and it associates a value with that key. Maps are unordered collections, meaning that the order of elements is not guaranteed.

## Creating a Map:

> var myMap map[keyType]valueType

Eg:
  
  ```
    var myMap = make(map[string]int)

```
  or
  
  
  ```
    my_Map:=map[string]int{
                           "Giri":3,
                           "Monish":7,
                           "Jagan":48
                        }
  ```
 
## Adding elements

```
 myMap[key] = value
 ```
  
  
## Deleting Elements



To delete an element from a map:

```
delete(myMap, key)
```
 the above line removes the element from map
 
## Checking if elements exist ?

```
package main
import ("fmt")

func main(){
	var mymap=make(map[string]string)
	mymap["name"]="Monish";
	mymap["interest"]="open source and devops"
	_,ok:=mymap["name"] // _ used in as variable to store the values when we are not going to use the value anymore,golang will not ask us to use this variable
	
	// _,check:=mymap["interest"]

	_,mon:=mymap["info"] 

	// here ok,mon are just varibles that hold boolean values 

	if ok{
		fmt.Println("value exist");
	}else{
		fmt.Println("value not exist for name field");
	}

	if mon{
		fmt.Println("the value exist for info");
	}else{
		fmt.Println("value not exist for key info");
	}

}
```


## Iterating over maps
  
   Maps are unordered, so you can't rely on a specific order when iterating. To iterate over a map, you can use a for range loop:
   
    ```
    for key, value := range myMap {
    // Process key and value
    }

	```
	
## Complex maps

  Maps with Complex Types:
  
   Maps can have keys and values of various types, including structs, slices, and other maps