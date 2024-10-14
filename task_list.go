package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TaskList struct {
	Tasks []Task
	Index int
}

func NewTaskList() TaskList {
	return TaskList{
		Tasks: []Task{},
		Index: 0,
	}
}

func (tl *TaskList) moveUp() {
	if len(tl.Tasks) > 0 {
		if tl.Index > 0 {
			tl.Index--
		}
	}
}

func (tl *TaskList) moveDown() {
	if len(tl.Tasks) > 0 {
		if tl.Index < (len(tl.Tasks) - 1) {
			tl.Index++
		}
	}
}

func (tl *TaskList) addDefault(timetick int) {
	tl.Tasks = append(tl.Tasks, Task{
		Title: fmt.Sprintf("Task @%d", timetick),
		Done:  false,
		OnPress: func() tea.Msg {
			return toggleCasingMsg{}
		},
	})
}

func (tl *TaskList) deleteSelected() {
	if len(tl.Tasks) == 0 {
		return
	}

	if tl.Index < len(tl.Tasks) && tl.Index >= 0 {
		tl.Tasks = append(tl.Tasks[:tl.Index], tl.Tasks[tl.Index+1:]...)
		if tl.Index > 0 {
			tl.Index--
		}
	}
}

func (tl *TaskList) markDone() {
	if len(tl.Tasks) == 0 {
		return
	}

	tl.Tasks[tl.Index].Done = !tl.Tasks[tl.Index].Done
}

func (tl *TaskList) selected() (*Task, bool) {
	if len(tl.Tasks) == 0 {
		return nil, false
	}

	task := &tl.Tasks[tl.Index]

	// FIX impossible condition
	if task == nil {
		return task, false
	}

	return task, true
}
