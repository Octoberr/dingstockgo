package tests

import "fmt"

type Phone interface {
	call(param int) string
	takePhoto()
}

type Huawei struct {
	name string
	price int
}

func (huawei *Huawei) call(param int) string {
	fmt.Printf("I am %s, I can call you %d\n", huawei.name, param)
	huawei.name = "swm"
	return huawei.name
}
func (huawei *Huawei) takePhoto()  {
	fmt.Println("I can take a photo for you")
}

func Exam1()  {
	var phone = &Huawei{}
	phone.name = "note3"
	phone.price = 122
	//phone := new(Huawei)
	phone.takePhoto()
	r := phone.call(50)
	fmt.Println(r)
	fmt.Println(phone.name)

}