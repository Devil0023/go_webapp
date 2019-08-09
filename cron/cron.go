package cron

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

//Setup 注册Crontab
func Setup() {
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
