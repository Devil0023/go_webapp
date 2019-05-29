package main

import (
	"../models"
	"fmt"
)

func main() {
	maps := make(map[string]interface{})
	maps["state"] = 1

	tag := models.GetTags(2, 20, maps)

	fmt.Println(tag)
}
