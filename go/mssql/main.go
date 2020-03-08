package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println(time.Now().Unix())
		time.Sleep(time.Minute)
	})
	c.Run()
}
