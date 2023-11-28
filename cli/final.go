package cli

import (
	"fmt"
	"log"
	"proyecto/models"

	tea "github.com/charmbracelet/bubbletea"
)

type finalModel struct {
	userID   string
	choices  []string
	detail   string
	price    uint
	cursor   int
	selected map[int]struct{}
}

func initialFinalModel(d string, p uint, userID string) finalModel {
	fm := finalModel{
		userID:   userID,
		choices:  []string{"Regresar al inicio", "Salir"},
		detail:   d,
		price:    p,
		cursor:   0,
		selected: make(map[int]struct{}),
	}
	// send the information to the database
	_, err := models.PlaceOrder(fm.userID, fm.detail, fm.price)
	if err != nil {
		log.Fatal(err)
	}
	return fm
}

func (m finalModel) Init() tea.Cmd {
	return nil
}

func (m finalModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			if m.cursor == 0 {
				om := initialOptionModel(m.userID)
				return om, om.Init()
			}
			if m.cursor == 1 {
				return m, tea.Quit
			}
			log.Fatal("out of bound option")
		}
	}
	return m, nil
}

func (m finalModel) View() string {
	s := "¡Maravilloso! Tu orden ha sido creada. Selecciona si deseas seguir en la app, o salir:\n\n"
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
