package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var Conf string

	info, err := ioutil.ReadFile("./.env")

	if err != nil {
		log.Fatal("Failed to parse env: ", err)
	}

	Env := string(info[:])

	//Conf := os.Stat("conf/app." + Env + ".ini")? "conf/app." + Env + ".ini": "conf/app.ini"

	_, err = os.Stat("conf/app." + Env + ".ini")

	if err != nil {
		Conf = "conf/app.ini"
	} else {
		Conf = "conf/app." + Env + ".ini"
	}

	fmt.Print(Conf)

}
