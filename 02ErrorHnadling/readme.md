# Error handling in go

## In golang we do not have try catch blocks where we have
     _,err:=command

    
 eg
     _,err:=os.Open(name);

	if err!=nil{
		return fmt.Errorf("cannot open the file the error states ///:%s ////",err);
	}

     create a error using 
```errors.New("this is used to create a simple custom error)   // use errors package 
    fmt.Errof("create a formatted error %s",err);
```
# How to handle errors in go using panic and recover functions 

## Panic() used to terminate the program when we have no way to handle the errors

 Recover() is a function that helps us to capture the value of panic , "handle the panic " The recover function is used to regain control of a panicking goroutine. It is only useful when called from within a deferred function. When a recover is called, it stops the propagation of the panic and returns the value that was passed to the panic. This allows you to handle the panic gracefully and perform cleanup or other necessary actions.

### recover() should always need to be used within defer functions 

what is defer; it is simple that a statement or a function followed by defer is executed only at the end of the program

```
         func open()(s string){

	_,err:=os.Open("file1.txt")
	defer func(){
          p:=recover();
		  if p!=nil{
			s="cannot open the file"
		  }
	}()
	if err!=nil{
		panic("error cannot open the file")
	}
	return s
}  
```


panic will never look the remaining part of the program except defer when it is executed ...;

recover captures the panic for eg
  defer func(){
     p:=recover()
     fmt.Println(p); // prints --hello monish there happened an error--
  }

  panic("hello monish there happened an error")



# Why panic
In Go, the panic and recover functions are typically used in exceptional cases where something unexpected or unrecoverable happens during program execution. While it's generally discouraged to use panic and recover for regular error handling, they serve essential roles in specific scenarios.

Handling Unrecoverable Errors:
panic is used to signal an unrecoverable error or a programming mistake. For example, if a critical data structure is in an invalid state or if a mandatory configuration file is missing, it might be appropriate to use panic. This helps prevent the program from continuing with corrupt or inconsistent data, which could lead to even more severe issues down the line.

Unexpected Situations in Critical Code Paths:
In critical sections of code where failure might have catastrophic consequences, you can use panic to ensure that such situations are immediately caught and handled. This is particularly useful in cases where you cannot recover from the error gracefully. An example might be an unrecoverable failure in a low-level system call.

Cleanup Operations:
panic can also be used to simplify cleanup operations in specific situations. When combined with the defer statement, panic can ensure that cleanup code is executed even if the program terminates abruptly due to a panic.


# Comparison

## Normal Error Handling:

Purpose: Normal error handling is the preferred way to handle errors in Go. It involves returning error values from functions and explicitly checking those error values to handle errors gracefully.
Use Case: Normal error handling is suitable for expected errors or situations where you can reasonably anticipate the error conditions and want to handle them gracefully without terminating the program. This is the recommended approach for most error scenarios in Go.
Advantages:
It provides a clear and controlled flow of error handling.
It enables the program to handle errors gracefully and continue functioning in a predictable manner.
It allows for better separation of error handling logic from the core business logic, leading to more maintainable code.
Disadvantages:
It requires developers to diligently check error values after each function call, which can lead to more boilerplate code.
The error handling code can sometimes obscure the main logic, reducing code readability.

## Panic/Recover Mechanism:
Purpose: The panic and recover mechanism is meant for handling exceptional, unexpected, and unrecoverable situations. It is generally used in cases where the program state is compromised, and there's no sensible way to continue the program's execution safely.
Use Case: Panic and recover should only be used in exceptional situations, such as critical programming errors, unrecoverable states, or situations where the program cannot safely recover and continue execution.
Advantages:
It simplifies error handling in certain critical code paths where recovery might not be possible or meaningful.
It helps prevent the program from continuing with corrupted data or inconsistent states, avoiding potentially catastrophic consequences.
It allows for cleanup operations to be performed even if the program terminates abruptly due to a panic, when combined with defer.
Disadvantages:
Overusing panic/recover can lead to less predictable program behavior and make debugging more challenging.
It should not be used as a substitute for regular error handling. Inappropriate use of panic/recover can make the program less reliable and more difficult to maintain.
It might mask underlying issues in the code if used as a band-aid solution for regular error handling.

#  Summary

>error in go is a type itself, errors is a package to use error functions and its operations 
    panic is used to teriminate ,recover should need to be used in defer functions only and it captures the value of panic 
    good practise mostly panic and recover should not be used to handle errors..
 
### Panic : 
   it is used only when we cannot handle that errors and has to terminate the code immediately, that is my code will not work without that part.
 
### Panic & Recover:

 when we put the approapriate Recover function in defer ,then the part where panic is encountered is come out of it and look for recover functions ,if recover is there then the remaining part of the code is executed and if not then golang terminates that execution flow.
 
   Generally recover is used in the case if we want to do some clean up process and exit the code.



