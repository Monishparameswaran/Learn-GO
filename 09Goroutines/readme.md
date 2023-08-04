# Goroutines

 "it is simply a small units that do incredible work", it would be more better for me if you know threads in java or any other programming language.we can say go has something similar to thread but with additional functionalities and lightweight feature

 ## What is a thread, in fact, before all that?
Threading is a process of separating the flow of execution within a process into multiple threads. Each thread represents an independent sequence of instructions, allowing for concurrent execution of tasks. Threads share the same memory space and system resources, enabling them to perform their tasks independently and concurrently.


in laymenn terms, Threading is a process of separating flow of execution of process into mutliple threads,where each thread has multiple tasks where the tasks are got executed independently 

By separating a process into multiple threads we can execute each thread independently and make the execution faster,since the threads are executed parallely it takes less time to complete the process.

>eg,
 If you have a program with functions for addition, subtraction, multiplication, and division that are executed sequentially one after another, there may be situations where some functions need to wait for others to complete. This can lead to inefficiency if those functions don't have any dependencies on each other.

By utilizing threading and executing these functions in separate threads, you can potentially achieve parallelism and make use of the CPU power more efficiently. Each function can run concurrently in its own thread, allowing them to execute simultaneously and independently. This can lead to improved performance, especially if the functions are computationally intensive or have long execution times.

Threading allows for concurrent execution, where multiple tasks or functions can be executed at the same time, taking advantage of the available CPU resources. 

# Go routine vs Threads

>***Threads are based on the operating system's threading model***, where the operating system schedules and manages the execution of threads. ***Goroutines, on the other hand, are part of the Go programming language and are managed by the Go runtime***. Goroutines have their own scheduling mechanism, which is optimized for concurrent execution.

> ***Goroutines are lightweight compared to threads***. While threads typically have a larger memory footprint and require more system resources, goroutines have a smaller memory footprint. This allows for the creation of a large number of goroutines without much overhead, making them more practical for concurrent programming.

> ***Goroutines in Go are designed to communicate and synchronize using channels. Channels provide a safe and efficient way for goroutines to share data and communicate with each other***. Goroutines can send and receive data through channels, allowing for easy coordination and synchronization between concurrent operations.

Goroutines enable concurrent programming, but they do not necessarily guarantee parallel execution. While goroutines can be executed in parallel if the Go runtime schedules them on different CPU cores, the actual parallelism depends on the available CPU resources and the scheduling decisions made by the Go runtime.
parallelism is not possible where go runtime manages thread execution by shifting the CPU usage by the threads ,continuos shifting based on time shows that they are executing concurrently but internally it does not.

>note a single CPU core can execute one instruction at a time

## How to create a go routine

```
package main

import (
	"fmt"
	"time"
)


func delayPrint(number int){

	time.Sleep(time.Second * 6);
	fmt.Println(number);
	
}
func main(){
	fmt.Println("number hello monish");

	for i:=0;i<=10;i++{
		go delayPrint(i); // this is goroutine "that function"
	}

	time.Sleep(time.Second * 40);
}


```

you all know in Java, main function is also a thread  similarly func main() is a goroutine ("Thread")

when the function  delayPrint() is called with go keyword then it becomes go routine ,this makes that function as independent since there is delay in each call which causes delay but the program ends before the complete execution of program this is because func main itself a go rountine which is indepent.

when you call go delayPrint(), it starts executing the delayPrint() function concurrently in a new goroutine. However, the main goroutine, which includes the main() function, continues its execution without waiting for the delayPrint() function to complete.since go routines are independent right,

In Go, when you invoke a function with the go keyword, it creates a new goroutine, which is a lightweight concurrent execution unit. The function specified after go is then executed concurrently, independently of the main goroutine.


# What is channel

> channels are a fundamental construct for communication and synchronization between goroutines. They provide a safe and efficient way to pass data and control between concurrent goroutines.

A channel is a pipeline through which you can send and receive values of a specific type. It acts as a communication medium, allowing goroutines to exchange data and coordinate their execution. Channels can be used to establish communication patterns like sharing data between goroutines, signaling events, and coordinating concurrent operations.

```
 func main() {
    // Create an unbuffered channel of type string
    ch := make(chan string)  // create a channel

    // Start a goroutine that sends a message through the channel
    go func() {
        ch <- "Hello, monish "
    }()   // anonymous function

    // Receive the message from the channel
    msg := <-ch
    fmt.Println(msg)
}

```

Channels can also be used for synchronization purposes. For example, you can use channels to ensure that certain goroutines wait for specific events to occur before proceeding. By blocking on channel operations, goroutines can synchronize their execution and coordinate their actions.

Overall, channels are a powerful feature in Go for facilitating communication and synchronization between goroutines, enabling safe and efficient concurrent programming.


# Is Concurrency and parallel processing are same ?


No , When a function call is precedes with a keyword called "go" then it becomes independent where usually when a func main() or main function in any programming language made a call then the main function has to wait for the function call to execute and return. 
      
Where preceding the function call with ***go*** violates that rule of waiting and make both independent of each other,such that func main will be executing in its own flow and the function call will be on some other universe.where both has no relation


***concurrency refers to the ability to execute multiple tasks independently and make progress on more than one task at a time.***

***
Parallelism, on the other hand, refers to the simultaneous execution of multiple tasks or processes. It involves the actual execution of tasks in parallel, where multiple processors, cores, or threads are utilized to perform different tasks simultaneously. 
 ***
 
 ***Concurrency focuses on managing multiple tasks independently and making progress on each task over time, whereas parallelism involves the simultaneous execution of multiple tasks.***
 
 - ***Concurrency*** does not necessarily require multiple processors or cores. It can be achieved even on a single processor by interleaving the execution of tasks. 
 
 -***Parallelism***, however, requires multiple resources (processors, cores, or threads) to execute tasks simultaneously.
 
 
>summary, while making functions independent of each other can be a factor in achieving concurrency, it does not fully define or encompass the concept of concurrency itself. Concurrency involves managing the execution of multiple tasks independently and concurrently, enabling efficient resource utilization and improved program performance.


# At last how a goroutine can be understood  ?
think there are two functions which has a powerful feature running independently , merely when the program start to run,both function call stand parallel to each other and the communication between them can be made by channels

```
func main(){

	Value:=make(chan int)
	go give(Value);  
	go get(Value);


	// both go routines run concurrently since channel is used in between them to communicate the value ,what happens here is
	// consider two buses going in the road where the people from one bus is getting to another through the pipe on the go 
	time.Sleep(time.Second *30);
}
func give(out chan int){
	for i:=0;i<=10;i++{
		time.Sleep(time.Second *2);
		out<-i;
	}
	close(out)
}
func get(in chan int){
	for val:=range in{
		fmt.Println(val);
	}
}
```
here give is someone who will give the value and get is someone who get and print the value ,dont you think it is a nightmare in case of other programming languages .
here since both functions are indepent and are standing in parallel to them , a pipe is connected between them such that the data transfer between functions can be done. ***the function that writes or sends data to the channel must close the channel after the workdone.***
>close(out)

only then the other end can stop recieving the data and do the remaining job.otherwise it could cause dead lock conditions.

# -Projects

### Write a program that creates two goroutines. Each goroutine should print a series of numbers from 1 to 5, with a slight delay between each number. Use the go keyword to launch the goroutines and observe the concurrent execution [solution]()

### Implement a simple producer-consumer scenario using goroutines and channels. Create a producer goroutine that generates random numbers and sends them to a channel. Create a consumer goroutine that receives these numbers from the channel and prints them.[solution]()

### Create a ticket generator. where your program should need to get the user details and store it in a struct ,use go routines and go channels to facilitate the ticket generation where it should need to be in random order 

# Select Statements

Formal definition
>The select statement in Go provides a way to handle multiple channel operations concurrently. It allows you to write code that can wait for communication on multiple channels simultaneously without blocking.


Lets break it down

consider a scenario where you wrote a function that acts as a intermediate , its job is simple,just get the data from multiple channels and send it to the printer function.
 
 the problem occurs here is the channel output are send in a sequence that you defined, not in the order " which comes first " . since,for an effiecient concurrency we need the order ***which comes first***
 
 ```
 func merger(in chan int,inMerge chan int,out chan int){
		if a,ok:=<-in{
		out<-a;
	}
	if b,ok:=<-inMerge{
	 out<-b;
	 }
	 close(out)
	 }
	
```
 
 here you can see there is a order where channel "in" has to send the data first the data ,only then data of channel 'inMerge'  can be send eventhough the channel inMerge returns the data first.
 
 To overcome this and forward the data as soon as it is returned,we go for select statements.
 As you guess, the select statements are used in this case.
 
 >syntax
 ```
 select {
case channelOperation1:
    // Code block executed when channelOperation1 is ready
case channelOperation2:
    // Code block executed when channelOperation2 is ready
...
case channelOperationN:
    // Code block executed when channelOperationN is ready
default:
    // Code block executed when none of the channel operations are ready
}
```


>ex
```
func merger(in chan int,inMerge chan int,out chan int){
		i:=0
		for {
			select{
			case a,ok:=<-in:   // check whether the channel return anythin //remember anything (incl nil )
				if ok{    // if it is not nil send the channel data to another channel
					out<-a;
				}else{
					i++;
				}
		case b,ok:=<-inMerge:
			if ok{
				out<-b;
			}else{
				i++;
			}
		}
		if(i>=2){
			break;
		}
		}
		close(out)
	}
	
	
	func print(out chan int,done chan bool){
		for val := range out{
			fmt.Println(val);
		}
		done<-true;
	}
```

### How it works


#### In the merger function:

- The select statement starts a loop that continuously monitors the two channel operations: <-in and <-inMerge.


- Whenever either of these channel operations becomes ready, the corresponding case block is executed.
- The first case <-in receives a value from the in channel.

- If the receive operation is successful (ok is true), the received value a is sent to the out channel using out <- a.

- If the receive operation on in fails (indicating the channel is closed), the ok value is false, and the else block is executed, incrementing the i counter.


- Similarly, the second case <-inMerge receives a value from the inMerge channel. If successful, the received value b is sent to the out channel.

- If the receive operation fails, the else block is executed, incrementing the i counter.
- After executing each case block, the program checks if i is greater than or equal to 2. If so, it breaks out of the loop.

Finally, the out channel is closed using close(out).

```
 
 ### what does it do
 
 - select statements has multiple cases like switch statements it select and execute the case which recieves the data at first from the channel irrespective of their order.
 
 
 
 
 >syntax
```.
 select{
case <-channel1:
    // Code to be executed when data is received from channel1
case data := <-channel2:
    // Code to be executed when data is received from channel2
case channel3 <- value:
    // Code to be executed when data is sent to channel3
default:
    // Code to be executed when no channel operation is ready
}
```

## What is buffered channels

 ***consider the scenario*** where there are 4 goroutines , where one goroutine "Generator" does the job of generating the numbers sequencially like 1,2,3,4,5,6,7,8...., "Splitter" is the goroutine that splits based on the even and odd .
 the odd values can be printed directly by the printer go routine but the even number need to be processed as its squares and has to be printed, here it takes 10 sec delay.
 
  Always remember one goroutine can process only one data at a time, when "Generator" generates a number,it has to be transfered to "Splitter" using channel , it do transfer only when there is no data is being processed in the reciever channel i.e "Splitter" here . now consider the number 2 is being waiting in the goroutine "Squarer" which was send by "Splitter", now when 4 enters "Splitter" it process and say 4 is even and that has to be sent to "Squarer" but 2 is already is being processed in the goroutine "Squarer", since one goroutine can process only one data at a time , "Squarer" cannot able to recieve the data , so data 4 will be waiting in "Splitter" ,since 4 is waiting in the goroutine "Splitter" , "Generator" cannot able to sent the data 5 to "Splitter",so 5 waits at "Generator" , 4 waits at "Splitter" and 2 is waiting at "Squarer" .
  
   just only because of data 2 all other data till end got affected and waiting.so to avoid this delay we introduce something called buffers ,nothing but a place to wait not in the goroutine but in the channel.
   
   imagine like a pipe has  sitting area to wait ,so that the goroutine can process its data without waiting for anyone.here the channel that connecting "Splitter and Squarer is made as buffered channel, so that the data 4 can wait in the sitting area "buffer" in channel ,so that each goroutine can process one data as per the flow without waiting for anyone.
   
   Where this scenario becomes life saver when there are 100000 data to be generated,splitted and printed. do you think waiting 10 sec per data is effiecient process ?

## Formal definitions and examples

In Go, a buffered channel is a channel with a fixed capacity to hold a certain number of values. Unlike an unbuffered channel, which blocks the sender until the receiver is ready, a buffered channel allows the sender to continue sending values as long as the channel is not full. The receiver can also read values from the channel at its own pace.

***To create a buffered channel in Go, you can use the make function with a second argument specifying the buffer size:***
```
channel := make(chan int, bufferSize)
```

```
package main

import (
	"fmt"
	"time"
)

func sender(channel chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending", i)
		channel <- i
	}
}

func receiver(channel <-chan int) {
	for i := 1; i <= 5; i++ {
		value := <-channel
		fmt.Println("Received", value)
		time.Sleep(1 * time.Second) // Simulating some processing time
	}
}

func main() {
	// Unbuffered Channel
	unbuffered := make(chan int)

	// Buffered Channel
	buffered := make(chan int, 3)

	go sender(unbuffered)
	go sender(buffered)

	go receiver(unbuffered)
	go receiver(buffered)

	time.Sleep(10 * time.Second) // Allowing goroutines to complete
}

```

In this example, we have a sender and a receiver function. The sender sends values from 1 to 5 to both an unbuffered channel (unbuffered) and a buffered channel (buffered). The receiver receives the values from both channels but simulates some processing time using time.Sleep(1 * time.Second).

If you run the code, you will observe the following behavior:

With the unbuffered channel: The sender will send a value, and the receiver will immediately receive and process it. The sender will be blocked until the receiver processes the value. This results in synchronous communication between the sender and receiver.

With the buffered channel: The sender will send all five values without getting blocked since the buffer size is set to 3. The receiver will start receiving and processing the values at its own pace. The sender will only block when the buffer is full. This allows for asynchronous communication, where the sender can continue sending values even if the receiver is slower.

By using buffered channels, you can decouple the sender and receiver, allowing them to work at different speeds and reducing the potential for goroutines to get stuck due to synchronization issues. Buffered channels are useful in scenarios where the sender and receiver operate at different rates or when a burst of values needs to be sent without immediate processing.


## WaitGroups

consider this with a scenario where you have called created multiple go routines in your func main , do you think that your main function will for the execution of that routines . 
  the answer is no,it will not, since go routines are independent ,func main will not wait for the result and terminate when its statements are over.
  to make the func main to wait for the goroutine we use something called ***waitgroups***

>#### Formal Definition

A sync.WaitGroup is a synchronization primitive provided by the Go standard library that allows you to wait for a collection of goroutines to complete their execution before proceeding with further code execution. It provides a simple way to coordinate the synchronization of goroutines.

The purpose of a sync.WaitGroup is to wait for a specified number of goroutines to call the Done method, indicating that they have completed their execution. Once the specified number of Done calls is reached, any goroutine waiting on the Wait method will be unblocked, allowing the program to continue its execution.

>eg
```
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup   // create a object to the waitGr

	numWorkers := 3
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // at this line go func main will wait for all the go routines to execute 
	// if you people remember previously we are using <-done or time.Sleep()

	fmt.Println("All workers completed")
}

```


#### You may ask why shouldnt we go for time.Sleep or <-done ?

    
Using time.Sleep or <-done instead of sync.WaitGroup to wait for goroutines to complete can lead to less efficient and less reliable code. 

***Explicit synchronization*** : sync.WaitGroup provides explicit synchronization and coordination of goroutines. It allows you to clearly express your intention to wait for specific goroutines to complete before proceeding with further code execution. This makes the code more readable and easier to understand for other developers.

***Dynamic number of goroutines***: If you have a dynamic number of goroutines, using sync.WaitGroup allows you to increment and decrement the counter as needed, without having to modify the waiting logic. This flexibility is not easily achieved with time.Sleep or <-done approaches.

***Efficient resource utilization***: sync.WaitGroup allows goroutines to complete and signal their completion explicitly through the Done method. This allows the Go scheduler to efficiently schedule other goroutines to execute while waiting, instead of wasting CPU cycles through arbitrary time.Sleep delays.

In summary, sync.WaitGroup is designed specifically for coordinating and synchronizing the completion of goroutines. It provides explicit and efficient synchronization, handles dynamic numbers of goroutines, allows for graceful shutdown, and facilitates error handling. Using time.Sleep or <-done as alternatives can result in less efficient and less reliable code, as they lack these benefits and may introduce unnecessary delays or race conditions