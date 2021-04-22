package tests

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}

}
func say2(s string)  {
	for i := 0; i<5; i++{
		time.Sleep(150*time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	//ch <- 0
	//close(ch)
}

func Fibonacci(n int, c chan int)  {
	x, y := 0,1
	for i :=0; i<n;i++{
		c <- x
		x, y = y, x+y
		time.Sleep(1000*time.Millisecond)

	}
	close(c)
}

func Factorial(n uint64) (result uint64)  {
	if (n>0){
		result = n*Factorial(n-1)
		return result
	}
	return 1
}

func Fibonacci1(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci1(n-2) + Fibonacci1(n-1)
}

func MapUse(){
	var CCM map[string]string
	CCM = make(map[string]string)
	CCM["1"] = "sadsadas"
	CCM["2"] = "sadhsakjdh"
}
