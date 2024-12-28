package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type addModel struct {
	form *huh.Form
}

var newTask Task
var add bool

func initAddModel() addModel {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What task do you want to add?").
				Value(&newTask.Task),
			huh.NewSelect[string]().
				Title("What status does the task have?").
				Options(
					huh.NewOption("Task is yet to be started", "open"),
					huh.NewOption("Currently working on this task", "active"),
					huh.NewOption("Task is finished", "done"),
				).
				Value(&newTask.Status),
			huh.NewConfirm().Title("Add this task?").Value(&add),
		),
	)
	return addModel{form: form}
}

func (m addModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m addModel) View() string {
	return m.form.View()
}

func (m addModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	if m.form.State == huh.StateCompleted {
		newTask.Date = time.Now().Format(time.DateOnly)
		oldTasks, err := ReadTaskDBFile("tasks.json")
		if err != nil {
			fmt.Printf("could not read all tasks form DB file. Error: %s\n", err.Error())
			os.Exit(1)
		}
		allTasks := append(oldTasks, newTask)
		err = SaveTaskDBFile("tasks.json", allTasks)
		if err != nil {
			fmt.Printf("could not add task to the DB file. Error: %s\n", err.Error())
			os.Exit(1)
		}
		// cmds = append(cmds, func() tea.Msg { return switchToListModel{} })
	}
	return m, tea.Batch(cmds...)
}
