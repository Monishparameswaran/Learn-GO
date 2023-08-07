package main

import (
    "fmt"
)

func main() {
    // Creating and initializing the map
    var myMap = make(map[string]int)

    myMap["Monish"] = 7
    myMap["Giri"] = 3

    fmt.Println(myMap)  // map[Giri:3 Monish:7]
}
