package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
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
			m.taskList.addDefault()
		case "delete":
		}
	}
	return m, nil
}

func (m model) View() string {
	result := fmt.Sprintf("Timetick: %d\n", m.timetick)
	result += fmt.Sprintf("TaskList: %v\n", m.taskList)
	return result
}
