package main
import "fmt"
import "sync"
import "time"
type Account struct{
	name string;
	money int;
	sync.Mutex
}
var mutex sync.Mutex;
var wg sync.WaitGroup
var DataBase []*Account;
func (obj *Account) checkBalance()int{
	return obj.money;
}
func (obj *Account) deposit(amount int){
	if(obj.TryLock() && amount <=100000){
	obj.money+=amount;
	obj.Unlock()
	obj.message(2,"Successfully Deposited..!")
	}else{
		go obj.message(2,"unsucessful")
	}
	
}
func (obj *Account)withDraw(amount int){
	if(obj.money>amount){
		obj.Lock()
		obj.money-=amount;
		obj.Unlock()
		fmt.Println("withdraw ",obj.name);
		obj.message(3,"Successful withDraw")
		}else{
			fmt.Println("not withdrawn ",obj.name);
			go obj.message(3,"unsucessful")
		}
}
func (obj *Account) message(num int,msg string){

	time.Sleep(time.Second * 5)
	fmt.Println();
	fmt.Println("--------------------Summary----------------------------");
	fmt.Println("user: ",obj.name);
	fmt.Println("balance: ",obj.money);
	switch(num){
	case 2:
		fmt.Println("Request: Deposit");
		fmt.Println("Status: ",msg);
	case 3:
		fmt.Println("Request: WithDraw");
		fmt.Println("Status: ",msg);
	}
	fmt.Println("-------------------------------------------------------");

	time.Sleep(time.Second * 1);
}

func getName() string{
	mutex.Lock();
	defer mutex.Unlock()
	fmt.Println("enter name:" );
	var name string;
	fmt.Scan(&name);
	return name;
}

func getNum()int{
	mutex.Lock()
	defer mutex.Unlock()
	var temp int;
	fmt.Scan(&temp)
	return temp;
}

func route(num int,name string,obj *Account){
	switch(num){
	case 1:
		go obj.message(0,"balance");
	case 2:
		time.Sleep(time.Second * 20)
		var money int;
		fmt.Println("Welcome, ",name,", please enter money to deposit");
		money=getNum();
		go obj.deposit(money);
	case 3:
		time.Sleep(time.Second * 20)
		var money int;
		fmt.Println("Welcome, ",name,", please enter money to withdraw");
		money=getNum()
		obj.withDraw(money)
	case 4:
		go createNewAccount(name);
	}
	wg.Done()
}
func createNewAccount(name string){
	DataBase = append(DataBase, NewConstructor(name))
	
		time.Sleep(time.Second * 6)
		fmt.Println();
		mutex.Lock()
		fmt.Println("created new account successfully....Welcome",name);
		time.Sleep(time.Second * 2)
		mutex.Unlock()
		
}
func NewConstructor(name string) *Account{
	return &Account{name:name};
}


func checkDataBase(name string)int{
	fmt.Println("checking Database....",name);
	var i int=-2;
	for index,usr:=range DataBase{
		if usr.name==name{
			i=index;
		}
	}
	time.Sleep(time.Second *2);
	return i;
}


func main(){
	
	for i:=0;i<10;i++{
		var name string;
		var num int;
		
		name=getName();
		val:=checkDataBase(name);
		if(val!=-2){
			fmt.Println("enter 1 to checkbalance");
			fmt.Println("enter 2 to Deposit");
			fmt.Println("enter 3 to withdraw");
			fmt.Println("enter __: ");
			fmt.Println("please wait, ",name,",we are processing your request")
			num=getNum()
			wg.Add(1)
			go route(num,name,DataBase[val])
		}else{
			fmt.Println(name," creating new Account...");
			time.Sleep(time.Second *2)
			createNewAccount(name)
		}
	}

	wg.Wait()
}