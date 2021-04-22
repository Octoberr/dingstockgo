package tests

import (
	"fmt"
	"spider/dbmange/mongo"
	"testing"
)

func Test_mongoConnect(t *testing.T) {
	//s1 := mongo.Student{Name: "小红", Age: 12}
	s2 := &mongo.Student{Name: "小兰", Age: 10}
	fmt.Println(s2)
	s3 := &mongo.Student{Name: "小黄", Age: 11}
	fmt.Println(s3)
	students := [] interface{}{*s2, *s3}
	fmt.Println(students)
	//mongo.MongoE(students)
}

