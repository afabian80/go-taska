package main

type model struct {
	timetick int
	taskList TaskList
}

func initialModel() model {
	return model{
		timetick: 0,
		taskList: NewTaskList(),
	}
}
