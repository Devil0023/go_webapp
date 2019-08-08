package main

import (
	"fmt"
	"go_webapp/pkg/gredis"
	"go_webapp/pkg/setting"
	"time"
)

func main() {

	setting.Setup()
	gredis.Setup()

	key := "ZL-TEST"
	value := "Hello World"

	duration := 3

	result, err := gredis.Set(key, value, duration)

	if err != nil {
		fmt.Println("here :", err)
	}

	fmt.Println("first set: ", result)

	result2, _ := gredis.Get(key)

	fmt.Println("secound get: ", string(result2))

	time.Sleep(5 * time.Second)

	result3, err2 := gredis.Get(key)

	fmt.Println("third get: ", string(result3), err2)
}
