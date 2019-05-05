package main

import (
	"fmt"
	"io/ioutil"
)

var env string

func main() {
	info, err := ioutil.ReadFile("./.env")

	if err != nil {
		fmt.Println("Can't Read Env: %s\n", err)
		panic(err)
	}

	env = string(info[:])

}
