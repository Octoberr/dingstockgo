package main

import "fmt"

func main(){
	var i int = 15
	fmt.Printf("%d的阶乘是%d\n", i, Factorial(uint64(i)))
	//var i int
	//for i=0; i<10; i++{
	//	fmt.Printf("%d\t", fibonacci(i))
	//}

}


func Factorial(n uint64)(result uint64){
	if (n>0){
		result = n*Factorial(n-1)
		return result
	}
	return 1
}


//斐波拉切数列
func fibonacci(n int)int{
	if n<2{
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}