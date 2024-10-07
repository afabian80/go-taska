package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	Title string
}

type model struct {
	timetick int
	state    string
	tasks    []Task
	index    int
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		state: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.state = "up"
			m.timetick++
			m.index = max(0, m.index-1)
		case "down":
			m.state = "down"
			m.index = min(m.index+1, len(m.tasks))
			m.timetick++
		case "a":
			m.timetick++
			m.tasks = append(m.tasks, Task{
				Title: fmt.Sprintf("Auto task at %d", m.timetick),
			})
		}
	}
	return m, nil
}

func (m model) View() string {
	result := ""

	for _, t := range m.tasks {
		result += fmt.Sprintf("task: %s\n", t.Title)
	}

	result += fmt.Sprintf("state: %s\n", m.state)
	result += fmt.Sprintf("tick: %d\n", m.timetick)
	result += fmt.Sprintf("index: %d\n", m.index)

	return result

}
