package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/deckarep/gosx-notifier"
)

func work(activityTime, restTime int) chan bool {
	var count int
	var notification string
	stop := make(chan bool)
	restTimer := new(time.Timer)
	activityTimer := new(time.Timer)
	go func() {
		activityTimer = time.NewTimer(5 * time.Second)
		count = 0
		for {
			select {
			case <-activityTimer.C:
				count++
				notification = fmt.Sprintf("You have completed pomodoro nr: %d\nTime to rest a little bit...", count)
				note := gosxnotifier.NewNotification(notification)
				note.Push()
				fmt.Println("Time to rest...")
				restTimer = time.NewTimer(3 * time.Second)
			case <-restTimer.C:
				note := gosxnotifier.NewNotification("Returning to work...")
				note.Push()
				activityTimer = time.NewTimer(5 * time.Second)
			case <-stop:
				return
			}
		}
	}()
	return stop
}

func main() {
	activityTime := flag.Int("time", 25, "Time to work")
	restTime := flag.Int("rest", 5, "Time to rest")

	flag.Parse()

	fmt.Println("Time: ", *activityTime)
	fmt.Println("Time: ", *restTime)

	stop := work(*activityTime, *restTime)
	time.Sleep(1 * time.Minute)
	stop <- true
}
