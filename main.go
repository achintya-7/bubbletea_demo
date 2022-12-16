package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on the list
	cursor   int              // which item our cursos is at
	selected map[int]struct{} // which items are selected
}

func initModel() model {
	return model{
		choices:  []string{"Github", "Twitter", "LinkedIn", "Portfolio"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil // we dont really wanna run anything at startup
}

func (m model) View() string {
	s := "What sites do you use? \n\n"

	for index, choice := range m.choices {
		cursor := " " // at starting cursor will be an empty string
		if m.cursor == index {
			cursor = "->"
		}

		// Is this choice selected
		checked := " "
		if _, ok := m.selected[index]; ok {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit"
	return s

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("\n\n\nBYE BYE!!!")
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
				fmt.Printf("You removed %s ", m.choices[m.cursor])
			} else {
				m.selected[m.cursor] = struct{}{}
				fmt.Printf("You have %s ", m.choices[m.cursor])
			}
		}
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
}
