package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.timetick++

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.taskList.moveUp()
		case "down":
			m.taskList.moveDown()
		case "a":
			m.taskList.addDefault(m.timetick)
		case " ":
			m.taskList.markDone()
		case "delete":
			m.taskList.deleteSelected()
		}
	}

	return m, nil
}

func (m model) View() string {
	result := fmt.Sprintf("Timetick: %d\n", m.timetick)
	result += fmt.Sprintf("TaskList: %v\n", m.taskList)

	return result
}
