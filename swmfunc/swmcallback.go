package swmfunc

import (
	"fmt"
	"os"
)

type cb func(int) int

func TestCallBack(x int, f cb) {
	f(x)
}

func CallBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

func WriteTofilr(ssdata string)  {
	dstFile,err := os.Create("product.json")
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer dstFile.Close()
	dstFile.WriteString(ssdata + "\n")
	fmt.Println(ssdata)
}