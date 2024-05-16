package shared

import "time"

func Scheduler() {
	NewTask("Garbage", "take out the garbage", 5)
	currTask := GetTask() // Gets current task obj

	time.AfterFunc(currTask.Rest, func() {
		SendNotif(currTask.Title)
	})

	select {}
}
