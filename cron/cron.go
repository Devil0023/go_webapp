package cron

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

//Run 运行Crontab
func Run() {
	log.Println("Start")

	c := cron.New()

	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Hello World by 5 secounds")
	})

	c.AddFunc("*/2 * * * * *", func() {
		log.Println("Hello World by 2 secounds")
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
