package cli

import (
	"fmt"
	"log"
	"proyecto/db"
	"proyecto/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"github.com/surrealdb/surrealdb.go"
)

type additionModel struct {
	userID   string
	choices  []models.Ingredient
	cursor   int
	dish     models.Dish
	selected map[int]struct{}
}

func initialAddModel(d models.Dish, userID string) additionModel {

	auxArr := make([]models.Ingredient, 0)

	for _, ing := range d.Additions {
		dbRes, err := db.DB.Select(ing)
		if err != nil {
			log.Fatal(err)
		}
		add := new(models.Ingredient)
		err = surrealdb.Unmarshal(dbRes, &add)
		if err != nil {
			log.Fatal(err)
		}
		auxArr = append(auxArr, *add)
	}

	am := additionModel{
		userID:   userID,
		choices:  auxArr,
		cursor:   0,
		dish:     d,
		selected: make(map[int]struct{}),
	}
	return am
}
func (m additionModel) Init() tea.Cmd {
	return nil
}
func (m additionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return nil, tea.Quit
		case "left":
			dm := initialStartModel(m.userID)
			return dm, dm.Init()
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			additionString := m.dish.Name + " ADICIONES: "
			var additionPrice uint
			for i := range m.selected {
				additionString += m.choices[i].Name + " "
				additionPrice += m.choices[i].Price
			}
			fm := initialFinalModel(additionString, m.dish.Price+additionPrice, m.userID)
			return fm, fm.Init()
		}
	}
	return m, nil
}
func (m additionModel) View() string {
	s := "\n\n" + m.dish.Detail + "\n\nPuedes elegir las siguientes adiciones:\n\nUsa la tecla espacio para marcar, cuando hayas terminado presiona Enter\n\n"
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}
		s += fmt.Sprintf("%s [%s] %s\t%d\n", cursor, checked, choice.Name, choice.Price)
	}
	s += "\nPress q to quit.\n"
	return wordwrap.String(s, 80)
}

func createOrder() {
	// We need dish name, additions and final price (dish price + additions price)

}
