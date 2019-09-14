package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

//Run 运行Crontab
func Run() {

	fmt.Println("Start")

	c := cron.New()

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Hello World by 5 seconds")
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
