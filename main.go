package main

func main() {
	taskList := TaskList()

	for _, task := range taskList {
		switch task.title {
		case "gym":
			SendNotif(task.message)
		}
	}
}
