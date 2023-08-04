	package main
	import ("fmt")
	import ("time")

	func main(){

		input:=make(chan int);
		inSquare:=make(chan int,100)
		inMerge:=make(chan int)
		in:=make(chan int)
		out:=make(chan int);
		done:=make(chan bool)
		doneSquare:=make(chan bool)
		go generator(input);
		go oddEvenSplitter(input,inSquare,inMerge,doneSquare)
		go squarer(inSquare,in)
		go merger(in,inMerge,out)
		go print(out,done)
		<-done
		<-doneSquare
	}

	func generator(input chan int){

		for i:=0;i<=100;i++{
			input<-i;
		}
		close(input)
	}
	func oddEvenSplitter(input chan int,inSquare chan int,inMerge chan int,doneSquare chan bool){

		for val:=range input{
			if(val%2==0){
				inSquare<-val
			}else{
				inMerge<-val;
			}
		}
		doneSquare<-true
		close(inSquare);
		close(inMerge);
	}
	func squarer(inSquare chan int,in chan int){
		for val:=range inSquare{
			time.Sleep(time.Second * 3)
			in<-val*val
		}
		close(in)
	}
	func merger(in chan int,inMerge chan int,out chan int){
		i:=0
		for {
			select{             /// select statements used forward the channel data that comes first 
			case a,ok:=<-in:
				if ok{
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