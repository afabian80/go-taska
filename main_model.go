package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type mainModel struct {
	timetick int
	taskList TaskList
}

func initialMainModel() mainModel {
	return mainModel{
		timetick: 0,
		taskList: NewTaskList(),
	}
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter", "return":
			t, ok := m.taskList.selected()
			if ok {
				return m, t.OnPress
			}
		}
	case toggleCasingMsg:
		return m.toggleSelectedItemCase(), nil
	}

	return m, nil
}

func (m mainModel) View() string {
	result := fmt.Sprintf("Timetick: %d\n", m.timetick)

	var cursor string

	selectedStyle := lipgloss.NewStyle().Background(lipgloss.Color("#00AA00"))

	for index, task := range m.taskList.Tasks {
		if index == m.taskList.Index {
			cursor = " > "
		} else {
			cursor = "   "
		}

		text := fmt.Sprintf("%2vTask: %v", cursor, task)
		if index == m.taskList.Index {
			result += fmt.Sprintln(selectedStyle.Render(text))
		} else {
			result += fmt.Sprintln(text)
		}
	}

	return result
}

func (m mainModel) toggleSelectedItemCase() tea.Model {
	task, ok := m.taskList.selected()
	if ok {
		lowerTitle := strings.ToLower(task.Title)
		upperTitle := strings.ToUpper(task.Title)

		if task.Title == upperTitle {
			task.Title = lowerTitle
		} else {
			task.Title = upperTitle
		}
	}

	return m
}
