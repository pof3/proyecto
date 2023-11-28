package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type optionModel struct {
	userID   string
	choices  []string
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialOptionModel(userID string) optionModel {
	m := optionModel{
		userID:   userID,
		choices:  []string{"Pedir comida", "Revisar mis órdenes"},
		selected: make(map[int]struct{}),
	}
	return m
}

func (m optionModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m optionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			// do as they wish
			// create order
			if m.cursor == 0 {
				sm := initialStartModel(m.userID)
				return sm, sm.Init()
			}
			if m.cursor == 1 {
				// check order option
				cm := initialCheckModel(m.userID)
				return cm, cm.Init()
			}
		}
	}
	return m, nil
}

func (m optionModel) View() string {
	s := "Te damos la bienvenida. A continuación, selecciona la acción que deseas realizar:\n"
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("\t%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}
