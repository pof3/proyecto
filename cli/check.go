package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"proyecto/models"
)

type checkModel struct {
	userID   string
	choices  []string
	cursor   int
	orders   []models.Order
	selected map[int]struct{}
}

func initialCheckModel(userID string) checkModel {
	// no error checking, the program will panic in the models function
	orders, _ := models.GetOrder(userID)

	cm := checkModel{
		userID:   userID,
		choices:  make([]string, 0),
		orders:   make([]models.Order, 0),
		cursor:   0,
		selected: nil,
	}

	cm.orders = orders

	return cm
}
func (m checkModel) Init() tea.Cmd {
	return nil
}
func (m checkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}
func (m checkModel) View() string {
	s := "\n\nAquí está tu lista de órdenes:\n\n"
	for _, choice := range m.orders {
		// Render the row
		s += fmt.Sprintf("Detalle: %s\t|\tPrecio: %d\t|\tFecha: %s\n", choice.Detail, choice.Price, choice.CreatedAt)
	}
	s += "\nPress q to quit.\n"
	return s
}
