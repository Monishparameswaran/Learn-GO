# Package ***welcome***
## what is a package in go?
    In go , any program starts with a name called package "name_of_the_package"
>packages in Go are a fundamental building block that promotes modularity, code reuse, encapsulation, 
dependency management, and code organization. They enable developers to create scalable, maintainable,
and efficient software systems.

## what are types of package in go
### Executable Packages:

An executable package is a package that contains a main function.
The main function serves as the entry point for the executable program.
When an executable package is compiled, it produces an executable binary that can be run independently.
Executable packages are typically used to create standalone applications or command-line tools.

### Library Packages:

A library package is a package that does not have a main function.
Library packages provide a collection of functions, types, constants, and variables that can be used by other packages or programs.
They are designed to be imported and used as dependencies in other projects.
Library packages encapsulate reusable code and provide modular functionality that can be leveraged by different applications.

# What is package main ?
>package main

When you specify package main in a Go program, it indicates that you are creating an executable program rather than a library package. The main package is a special package in Go that serves as the entry point for executable programs.

***In Go, a project should have only one main package.***

## Should i need to name a program as ***main.go***  in order to execute as main ?

 
No, you are not required to name your Go program file as main.go in order to execute it as the main package. The name of the Go program file can be anything you prefer, as long as it has the .go extension.

# Where is the package main

 there is no separate "main package" in Go. Instead, the package main declaration is used within a Go source file to indicate that it belongs to the main package and contains the main function.
 
 In simple, you cannot directly found the package main in go ,where it is so called package main where a group of program that starts with ```package main``` are part of main package , there by forming a package main.
 
## Consider i have mulitple program .go with package main but has func main in only in one program among all named with package main, will it work.

  that is in same folder
  
  >In main.go
  
  >```
package main
func det(){
	printinfo();
	user();
}	
  ```

In function.go
  

package main
import ("fmt")

func printfunc(){
	fmt.Println("hello monish")
}

```
  
  it will work , in simple in a project if there are multiple file.go with package main then they are belongs to a package called "main" and condition is ***there can be multiple .go with package main but only one func main() amongst all***
  
  no magic it just a logic ,that a project  should have only one starting point right? , that is our func main().
  
  > anything in a same folder is belongs to same package.any functions in one another program  that is in same package need not to be imported to work with its functionalities.for example consider the above program that main.go and function.go belongs to same package called as "main" here it is not needed to import printfunc() that is in function.go to the main.go to use that.
  
## How to run this
  >go run main.go function.go
  
  
  ***good to know***: a function starting with small letter is a private function by default
  
  ## How to run a go program
  >go run main.go
  ## How to run a go package
  >go run ./ 
  
  laymenn term ,whenever you see "go package" remember there are multiple go programs in it,
  ***remember you should be in the same directory only then it will work***
  # Creating go library and working with it.
  
  ### i) create,push to github and import
  
  >mkdir new_project
  
  >cd new_project
  
  >go mod init github.com/user_name/name_of_library
  
  
  create a file_name.go and write some code ,remember you should not write a func main() since it is  a library.
  
  eg
  ``` 
  
  package monish
 
  import "fmt"
  func Say(){
      fmt.Println("hello geeks...!");
      }
  func Thank(){
      fmt.Println("Thank you");
      }
 
 ```
 
 >git remote add origin "remote_repo_link"
 
 >git add .
 
 >git commit -m "initial commit"
 
 >git push origin github.com/user_name/name_of_library
 
 >git tag "v1.01"
 
 >git push v1.01
 
 now your library is exported to globe 
 
 ## How to import and use external Library
 
 now go to the folder where your project folder that contains main.go
 
 >go mod init example.com/user_name/myproject
 
 ```
  package main
  
  import "github.com/user_name/name_of_library"
  
  func main(){
   monish.Say(); // here package_name.function_name
   monish.Thank();
   
  }
  
  ```
  
 >go get -u github.com/user_name/name_of_library
 
 >go run ./
 
 all done , now you should able to see the output.
  # How to initialize a go workspace (go mod init)
  
  The go mod init command is used to initialize a new Go module in your project directory. It creates a go.mod file that tracks your project's dependencies and provides versioning information.

The go.mod file will be created in your project directory, and it serves as the central configuration file for your module. It includes information about your module's dependencies, their versions, and other relevant metadata.

After running go mod init, Go will automatically attempt to download the necessary dependencies based on the import statements in your Go source files. If any dependencies are missing, Go will fetch them and add them to the go.mod file.
 
 ## How to import and use local library
 
Let us assume you have your library created locally and not pushed to github and you have to import to your local project
eg

Directory name : Dhoni_directory

``` 
contents of go mod

module github.com/Monishparameswaran/libfile

go 1.20

```
```
package anything
import "fmt"

func PrintMonish(){
	fmt.Println("hello geeks and monish");
}
```
in library folder

>go build

#### lets create a local project somewhere

>go mod init example.com/user_name/projectname

```
package main
import (
	"github.com/Monishparameswaran/libfile/Dhoni_directory"
)
func main(){
	anything.PrintMonish();   //package.function
}
```

before running this code

>go mod edit -replace github.com/Monishparameswaran/libfile/Dhoni_directory=/home/mon/Dhoni_directory

#### what is this ? 
here we are telling go to edit go.mod file to change the library path to this real path.remember the module path is fake and does not exist and go expects module path and os expects the actual path to find the file . so we have to do this.

this ***/home/mon/Dhoni_directory*** is nothing but a ***actual path*** of the directory


 
 >#### In go we have no private or public keywords instead if the first letter of the  function or variable defined with small letter then it is private only accessible within the package but if it starts with capital then it is public .
 
 
 example 
 
    func Print()  // public
    
    func print()  //private
 
 
 
 
 ## Another type is your Library file is in same Workspace or same module
   
   Myproject
    
    |
    |--> Main
        |->main.go  # package main
    |--> Library
        |->function.go # package monish
    |--> go.mod
         module local-project
    
    
    IN main.go
     package main
     import ( "local-project/Library")
     func main(){
        monish.printHello();
     }
     




here

       local-project/Library
	      |              |
		  v              --> name of the folder where the file is in
	   (module name that is in go.mod)


 #  # what is init functions
 
 
init function is a special function that is automatically executed before the main function in a Go program. It is used for initialization purposes, such as setting up variables, configuring the environment, or registering resources.
 
 No Arguments or Return Values: The init function doesn't accept any arguments and doesn't return any values. Its purpose is to perform initialization tasks rather than producing a result.
 
 No Invocation: Unlike other functions, you don't need to explicitly call an init function. It is invoked automatically by the Go runtime