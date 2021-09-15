package main


import "fmt"

func main(){
	nums := []int{2, 3, 4}
	sum := 0
	for _, num :=range nums{
		sum += num
	}
	//在数组上使用range会传入index和值两个变量，不需要的index直接使用_省略，需要的时候可以将其打印
	fmt.Println("sum:", sum)

	//下面演示打印数组的索引
	for i, num := range nums{
		//if num == 3{
		//	fmt.Println("index:", i)
		fmt.Printf("%d -> %d\n", i, num)
		//}
	}

	//range用来自循环map的键值对
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	//用range来枚举Unicode字符串，同理第一个参数是字符索引， 第二个参数是字符Unicode本身
	for i, c := range "go"{
		fmt.Println(i, c)
	}

}