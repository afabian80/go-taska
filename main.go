package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type OptionalInt struct {
	value int
	ok    bool
}

type Task struct {
	Title string
}

type model struct {
	timetick int
	state    string
	tasks    []Task
	index    OptionalInt
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
		timetick: 0,
		state:    "",
		tasks:    []Task{},
		index:    OptionalInt{},
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
			if m.index.ok {
				m.index = OptionalInt{
					value: max(0, m.index.value-1),
					ok:    true,
				}
			}
		case "down":
			m.state = "down"
			m.timetick++
			if m.index.ok {
				m.index = OptionalInt{
					value: min(m.index.value+1, len(m.tasks)-1),
					ok:    true,
				}
			}
		case "a":
			m.timetick++
			m.tasks = append(m.tasks, Task{
				Title: fmt.Sprintf("Auto task at %d", m.timetick),
			})
			m.index.ok = true
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
	result += fmt.Sprintf("index: %v\n", m.index)

	return result

}
