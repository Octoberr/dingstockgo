package main

import (
	"fmt"
)
/*定义一个接口Phone，接口里面有一个方法叫call()。*/
type Phone interface{
	call()
}

type NokiaPhone struct{

}


func (nokiaPhone NokiaPhone) call(){
	fmt.Println("T am Nokia, I can call you!")
}

type IPhone struct{

}

func (iPhone IPhone) call(){
	fmt.Println("I am iPhone, I can call you!")
}

func main(){
	//定义一个Phone类型变量，分别赋值为NokiaPhone, IPhone,调用call()然后输出
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone=new(IPhone)
	phone.call()
}