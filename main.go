package main

import (
	"time"
)

func main() {
	list := TodoList()

	currentTime := time.Now()

	switch hour := currentTime.Hour(); hour {

	case 0:
		SendNotif(list[0])
	}

}
