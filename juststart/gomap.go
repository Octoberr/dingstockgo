package main

import "fmt"

func main() {
	var countryMap map[string]string
	countryMap = make(map[string]string)
	countryMap["France"] = "Pairs"
	countryMap["Italy"] = "Rome"
	countryMap["China"] = "Beijing"

	for country := range countryMap{
		fmt.Println("capital of", country,"is", countryMap[country])
	}
	/* 查看元素在集合中是否存在 */
	//captial, ok := countryMap["United States"]
	captial, ok := countryMap["China"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}

}
