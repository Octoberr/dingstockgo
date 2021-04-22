package tests

import "fmt"

type Person struct {
	age int
	name string
	car Car
}

type Car struct {
	name string
}

var personMap map[string]*Person

func setName(person *Person, name string)  {
	person.name = name
}

func (person *Person) setName(name string)  {
	person.name = name
}

func printName(person Person)  {
	fmt.Println(person.name)

}

func (person Person) printName()  {
	fmt.Println(person.name)

}

func Examp() {
	person := Person{}
	fmt.Println(person) //{0  {}}
	person.age = 12
	person.name = "小明"
	person.car = Car{"宝马"}
	fmt.Println(person) //打印{12 小明 {宝马}}
	setName(&person, "小红")
	fmt.Println(person) //{12 小红 {宝马}}, 成功赋值!
	person.setName("小黑")
	fmt.Println(person) //{12 小黑 {宝马}}, 成功赋值!
	personMap = make(map[string]*Person)
	personMap["test"] = &person
	person = *personMap["test"]
	person.name = "小兰"
	fmt.Println(person) //{12 小兰 {宝马}},成功赋值!
	for _, value := range personMap {
		fmt.Println(value) //&{12 小兰 {宝马}},读取的也是正确的小兰
	}
}