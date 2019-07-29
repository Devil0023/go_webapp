package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Start")

	c := cron.New()

	c.AddFunc("* * * * * *", func() {
		log.Println("Hello World")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 60)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
