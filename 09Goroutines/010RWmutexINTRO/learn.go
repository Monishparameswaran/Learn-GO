package main
import "fmt"
import "sync"

/*
      In go sync package provides two types of lock one is write lock and another one is write lock , where
	   i) sync.Mutex has only write locks
	   ii) sync.RWMutex has both read and write locks

	   Remember one thing , when is there is a write lock,no other locks are allowed to access the resource , it may be write lock or read lock 
	   but if there is a read lock ,then all other read locks are allowed to read the resource.
	   */
var count int=33
func read(){
	mutex:=sync.RWMutex{};
	mutex.RLock();
	fmt.Println("Is it possible to RLock--",mutex.TryRLock());// here you will get output as true where it is possible to RLock the resource that is RLocked previously
	fmt.Println("Is it possible to Lock--",mutex.TryLock()); // try lock is used in case where to check whether it is possible to lock the resource since locking directly could cause deadlock in that case
	if(mutex.TryRLock()){
		mutex.RLock();
		fmt.Println(count);
		mutex.RUnlock()
	}
	mutex.RUnlock()
}
func main(){
	read();
}