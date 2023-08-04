package main

import (
	"Nasa-space/Nasa"
)

/*
now initially it will not work because it is local and it is in another workspace , so we have to give
go mod edit -replace "Nasa-space/Nasa"=/home/monish/golang-projects/learn-go/01Packages/03LocalPackages2/Nasa
                                              |
                                              -> the actual file path of that files directory
*/

func main(){
	nasa.LaunchVechicle();
}