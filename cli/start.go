package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"proyecto/models"
)

type startModel struct {
	userID          string
	restaurantNames []string
	restaurants     []models.Restaurant
	cursor          int
	selected        map[int]struct{}
}

func initialStartModel(userID string) startModel {
	restaurantArray, _ := models.GetRestaurants()
	m := startModel{
		selected: make(map[int]struct{}),
	}
	for _, r := range restaurantArray {
		m.restaurantNames = append(m.restaurantNames, r.Name)
	}
	m.restaurants = restaurantArray
	m.userID = userID
	return m
}

func (m startModel) Init() tea.Cmd {
	return nil
}
func (m startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.restaurantNames)-1 {
				m.cursor++
			}
		case "enter", " ":
			dm := initialDishModel(m.restaurants[m.cursor], m.userID)
			return dm, dm.Init()
		}
	}
	return m, nil
}

func (m startModel) View() string {
	s := "A continuación encontrarás los restaurantes disponibles:\n"

	for i, choice := range m.restaurantNames {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("\t%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}
