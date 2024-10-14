package main

import tea "github.com/charmbracelet/bubbletea"

type Task struct {
	Title   string
	Done    bool
	OnPress func() tea.Msg
}

type toggleCasingMsg struct{}
