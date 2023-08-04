	package main
	import "fmt"
	import "sync"
	import "time"
	const end=100
	var num int=0
	// mutex acts like a locker of the memory location ,where deadlock may occur due to access of same memory location
	//to avoid this we use locker when one function use one resource then the resource is locked and unlocked only after the access
	func increment(done chan bool,mutex *sync.Mutex){
		for i:=0;i<=end;i++{
			time.Sleep(time.Second * 3)
			mutex.Lock()
			num++;
			mutex.Unlock()
			fmt.Println(1," ",i);
		}
		done<-true  // note here done channel is not closed,
						// here the main purpose is just to send a find a final signal,
						//not a data transfer between sender and reciever, here there is no one waiting for this done at reciever side
						// what about main function,see that waits for a single value right and movesoff only in case of serious of value 
						// generation its mandatory to close the channel
	
	}

	func decrement(done chan bool,mutex *sync.Mutex){
		
		defer fmt.Println("decrement channel executed")
		for i:=0;i<=end;i++{
			time.Sleep(time.Second * 3)
			mutex.Lock()
			num--;
			mutex.Unlock()
			fmt.Println(2," ",i);
		}
		
		
	}
	func main(){

		mutex:=&sync.Mutex{};
		done:=make(chan bool);
		go increment(done,mutex);
		go decrement(done,mutex);
		<-done
		<-done
		fmt.Println(num);

	}