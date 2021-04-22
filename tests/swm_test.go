package tests

import (
	"encoding/json"
	"fmt"
	"spider/dbmange/redis"
	"testing"
	"time"
)

var rdb redis.RdbConf

func init() {
	rdb.Address = "47.108.176.230"
	rdb.Port = "6379"
	rdb.DB = 0
	rdb.PassWD = "swm123"
}

func Test_Redis(t *testing.T) {
	res := redis.ExampleClient(&rdb)
	t.Log(res)
}

func Test_json(t *testing.T) {
	timeNow := time.Now().String()
	fmt.Println("Time now:", timeNow)
	if result, err := json.Marshal(&rdb); err == nil{
		fmt.Println(result)
		fmt.Println(string(result))
	}
}

func Test_TimeFormat(t *testing.T) {
	//时间格式化
	currentTime := time.Now()
	fmt.Println("当前时间  : ", currentTime)
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	//时间解析
	withNanos := "2006-01-02T15:04:05"
	t1, e := time.Parse(
		withNanos,
		"2012-11-01T22:08:41")
	if e == nil{
		fmt.Println(t1.Format("2006-01-02 15:04:05"))
	}
}

func Test_RedisPool(t *testing.T) {
	redisAdd := fmt.Sprintf("%s:%s", rdb.Address, rdb.Port)
	fmt.Println(redisAdd)
	pool := redis.PoolInitRedis(redisAdd, rdb.PassWD)
	c1 := pool.Get()
	c2:=pool.Get()
	c3:=pool.Get()
	c4:=pool.Get()
	c5:=pool.Get()
	fmt.Println(c1,c2,c3,c4,c5)
	time.Sleep(time.Second * 5)//redis一共有多少个连接？？
	_ = c1.Close()
	_ = c2.Close()
	_ = c3.Close()
	_ = c4.Close()
	_ = c5.Close()
	time.Sleep(time.Second*5) //redis一共有多少个连接？？

	//下次是怎么取出来的？？
	b1:=pool.Get()
	b2:=pool.Get()
	b3:=pool.Get()
	fmt.Println(b1,b2,b3)
	time.Sleep(time.Second*5)
	_ = b1.Close()
	_ = b2.Close()
	_ = b3.Close()

	//redis目前一共有多少个连接？？
	for{
		fmt.Println("主程序运行中....")
		time.Sleep(time.Second*1)
	}
}

func Test_addr(t *testing.T) {
	Examp()
}

func Test_interface(t *testing.T) {
	//ExpInterface()
	//Exam1()
	res := Md5("swm123")
	t.Log(res)
}

func Test_goroute(t *testing.T) {
	//ch := make(chan int)
	go say2("lq")
	say("swm")
	//fmt.Println(<-ch)
}

func Test_fac(t *testing.T) {
	i := 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

}
func Test_fb(t *testing.T) {
	for i :=0; i<10; i++{
		t.Logf("%d\t", Fibonacci1(i))
	}

}