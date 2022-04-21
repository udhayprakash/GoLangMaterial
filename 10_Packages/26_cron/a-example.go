package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron"
)

func crontsk() {
	c := cron.New()
	c.AddFunc("@every 5s", addstamp)
	c.AddFunc("@every 60s", cleardoc)
	c.Start()
	time.Sleep(time.Duration(1<<63 - 1)) // running for more than 260 years
}

func addstamp() {
	t := time.Now()
	t.String()
	crondoc(t.Format("2006-01-02T15-04-05.000Z"))
}

func crondoc(msg string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	file.WriteString(msg + "\n")
	defer file.Close()
}

func cleardoc() {
	if err := os.Truncate("log.txt", 0); err != nil {
		fmt.Println("Failed to truncate: %v", err)
	}
}

func main(){
	crondoc("Udhay")
	crontsk()
}