package main

type task struct {
	title   string
	message string
	rest    int
}

func newTask(title string, message string, rest *int) *task {
	// Default rest period if nil == 5
	defaultRest := 5
	if rest != nil {
		defaultRest = *rest
	}

	// Makes new task with inputed fields
	t := task{title: title, message: message, rest: defaultRest}

	return &t // return object
}

func setInt(number int) *int {
	x := number
	return &x
}

func TaskList() []*task {
	garbage := newTask("garbage", "Go take out the garabge, you have 5 minutes...", nil)
	gym := newTask("gym", "Go workout, you have up to 2 hours", setInt(120))
	taskList := []*task{garbage, gym}

	return taskList

}
