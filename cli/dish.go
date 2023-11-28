package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"proyecto/models"
)

type dishModel struct {
	userID     string
	choices    []string
	cursor     int
	restaurant models.Restaurant
	selected   map[int]struct{}
}

func initialDishModel(r models.Restaurant, userID string) dishModel {
	dm := dishModel{
		restaurant: r,
		selected:   make(map[int]struct{}),
	}
	choices := make([]string, 0)
	for _, d := range dm.restaurant.DishesFields {
		choices = append(choices, d.Name)
	}
	dm.userID = userID
	dm.choices = choices
	return dm
}

func (m dishModel) Init() tea.Cmd {
	return nil
}

func (m dishModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			sm := initialStartModel(m.userID)
			return sm, sm.Init()
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
			seed := m.restaurant.DishesFields[m.cursor]
			am := initialAddModel(seed, m.userID)
			return am, am.Init()
		}
	}
	return m, nil
}
func (m dishModel) View() string {
	s := "Los platos ofrecidos por este restaurante son:\n\n"
	for i, choice := range m.restaurant.DishesFields {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("\t%s %s\t%d\n", cursor, choice.Name, choice.Price)
	}
	s += "\nPress q to quit.\n"
	return s
}
