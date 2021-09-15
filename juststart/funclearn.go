package main

import (
	"fmt"
	"math"
)





func main(){
	var a int = 100
	var b int = 200
	var ret int

	x, y :=swap("swm", "lc")

	ret = max(a, b)

	fmt.Printf("最大值是：%d\n", ret)
	fmt.Println(x, y)

	/*生命函数变量， 这是一个匿名函数吧*/
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(25)
	}

	/*使用函数*/
	fmt.Println(getSquareRoot(9))

}


func max(num1, num2 int) int{
	var result int
	if (num1 > num2){
		result = num1
	}else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string){
	return y, x
}