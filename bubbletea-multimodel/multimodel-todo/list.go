package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type listModel struct {
	table table.Model
}

func initListModel() listModel {
	return listModel{
		table: generateTableFromJSON(),
	}
}

func generateTableFromJSON() table.Model {
	tasks, err := ReadTaskDBFile("tasks.json")
	if err != nil {
		fmt.Printf("could not read all tasks form DB file. Error: %s\n", err.Error())
		os.Exit(1)
	}
	columns := []table.Column{
		{Title: "Date", Width: 10},
		{Title: "Task", Width: 45},
		{Title: "Status", Width: 20},
	}
	var rows []table.Row
	for _, task := range tasks {
		r := table.Row{task.Date, task.Task, task.Status}
		rows = append(rows, r)
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("#669bbc")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) View() string {
	return m.table.View()
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Println("You have selected task:", m.table.SelectedRow()[1]),
			)
		case "a":
			return m, func() tea.Msg { return switchToAddModel{} }
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
