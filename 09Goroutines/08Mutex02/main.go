package main
import "fmt"
import "sync"

type Counter struct{
	counter int;
}
func main(){
	mutex:=sync.Mutex{};
    obj:=Counter{counter:0}
	done:=make(chan bool)
	go func(){
		for i:=0;i<=10000;i++{
			mutex.Lock()
			obj.counter++;
			mutex.Unlock()
		}
		done<-true
	}()

	go func(){
		for i:=0;i<=10000;i++{
			mutex.Lock();
			obj.counter--;
			mutex.Unlock();
		}
		done<-true
	}()

	<-done
	<-done

	fmt.Println(obj.counter);
	
}