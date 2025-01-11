package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	activeModel tea.Model
}

func initMainModel() mainModel {
	return mainModel{
		activeModel: initListModel(),
	}
}

func (m mainModel) View() string {
	return m.activeModel.View()
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

type switchToListModel struct{}
type switchToAddModel struct{}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case switchToAddModel:
		m.activeModel = initAddModel()
		return m, m.activeModel.Init()
	case switchToListModel:
		m.activeModel = initListModel()
		return m, m.activeModel.Init()
	}
	m.activeModel, cmd = m.activeModel.Update(msg)
	return m, cmd
}

func main() {
	p := tea.NewProgram(initMainModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
