package ui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Bhargav16exd/nginxctl/commands"
	"github.com/Bhargav16exd/nginxctl/generators"
)

// model - state of ui
type model struct {
	sitesAvailablePath               string
	sitesEnabledPath                 string
	choices                          []string
	choicesInfo                      []string
	cursor                           int
	homeComponentActive              bool
	resetComponentActive             bool
	generateApiConfigComponentActive bool
	selected                         map[int]struct{}
	focusIndex                       int
	inputs                           []textinput.Model
	quitting                         bool
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	focusedButton = focusedStyle.PaddingLeft(2).Render("< Yeah this is button [Enter] it to Generate Config >")
	blurredButton = fmt.Sprintf("%s", blurredStyle.PaddingLeft(2).Render("< Yeah this is button [Enter] it to generate Config >"))
)

func runNginxSetup() (string, string) {

	commands.CheckNginxInstallation()

	path, err := commands.FetchNginxConfPath()

	if err != nil {
		return "", ""
	}

	sitesAvailablePath, sitesEnabledPath := commands.CheckCreateSitesDir(path)

	return sitesAvailablePath, sitesEnabledPath

}

func InitialModel() model {

	sitesAvailablePath, sitesEnabledPath := runNginxSetup()

	m := model{

		sitesAvailablePath: sitesAvailablePath,
		sitesEnabledPath:   sitesEnabledPath,

		//Inputs
		inputs: make([]textinput.Model, 3),

		// Choices
		choices: []string{"Reset Existing Config ", "Genereate New Configration"},

		choicesInfo: []string{"[ Restore your current nginx.conf to factory defaults, before using the tool ]",
			"[ Create a new nginx configuration optimized for API servers with reverse proxy setup. ]"},

		homeComponentActive:              true,
		resetComponentActive:             false,
		generateApiConfigComponentActive: false,

		selected: make(map[int]struct{}),
	}

	var t textinput.Model

	for i := range m.inputs {

		t = textinput.New()
		t.CharLimit = 32
		t.SetWidth(30)

		s := t.Styles()
		s.Cursor.Color = lipgloss.Color("205")
		s.Focused.Prompt = focusedStyle
		s.Blurred.Prompt = blurredStyle

		s.Focused.Placeholder = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

		s.Blurred.Placeholder = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

		s.Blurred.Text = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED"))

		t.SetStyles(s)

		switch i {
		case 0:
			t.Placeholder = "Project Name"
			t.Focus()
		case 1:
			t.Placeholder = "Localhost Target Port"
		case 2:
			t.Placeholder = "Domain"
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "up":

			if m.cursor > 0 {
				m.cursor--
			}

			if m.generateApiConfigComponentActive {
				m.focusIndex--

				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs)
				}

				cmds := make([]tea.Cmd, len(m.inputs))

				for i := 0; i <= len(m.inputs)-1; i++ {
					if i == m.focusIndex {
						cmds[i] = m.inputs[i].Focus()
						continue
					}
					m.inputs[i].Blur()
				}

				return m, tea.Batch(cmds...)
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

			if m.generateApiConfigComponentActive {

				m.focusIndex++

				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs)
				}

				cmds := make([]tea.Cmd, len(m.inputs))

				for i := 0; i <= len(m.inputs)-1; i++ {
					if i == m.focusIndex {
						cmds[i] = m.inputs[i].Focus()
						continue
					}
					m.inputs[i].Blur()
				}
				return m, tea.Batch(cmds...)
			}

		case "esc":

			m.homeComponentActive = true
			m.resetComponentActive = false
			m.generateApiConfigComponentActive = false

		case "enter":

			if m.cursor == 0 && m.homeComponentActive {
				m.homeComponentActive = false
				m.resetComponentActive = true
				m.generateApiConfigComponentActive = false
				break
			}

			if m.cursor == 1 && m.homeComponentActive {
				m.homeComponentActive = false
				m.resetComponentActive = false
				m.generateApiConfigComponentActive = true
				break
			}

			if m.focusIndex == len(m.inputs) && m.generateApiConfigComponentActive {

				res := generators.GenerateApiConfig(m.inputs[0].Value(), m.sitesAvailablePath, m.inputs[1].Value(), m.inputs[2].Value())

				if res {
					m.homeComponentActive = true
					m.resetComponentActive = false
					m.generateApiConfigComponentActive = false
					break
				}

			}
		}

	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m model) View() tea.View {

	// header section
	s := Header + "\n\n"

	if m.homeComponentActive {

		s += BoxStyle.Render("Home") + "\n"

		for i, choice := range m.choices {

			cursor := " "
			if m.cursor == i {
				cursor = ">"
				s += ItemStyle.Render(fmt.Sprintf("\n%s  %s\n   %s\n", cursor, choice, m.choicesInfo[i]))
			} else {
				s += fmt.Sprintf("\n%s  %s\n   %s\n", cursor, choice, m.choicesInfo[i])
			}

		}
	}

	if m.resetComponentActive {
		s += BoxStyle.Render("Reseting Config")
	}

	if m.generateApiConfigComponentActive {

		s += BoxStyle.Render("Generate Config") + "\n"

		var b strings.Builder

		for i := range m.inputs {

			b.WriteString(InputBoxStyle.Render(m.inputs[i].View()))

			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton

		if m.focusIndex == len(m.inputs) {
			button = &focusedButton
		}

		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

		if m.quitting {
			b.WriteRune('\n')
		}

		s += b.String()
	}

	// The footer
	s += "\n\n\n\n\n----------------------------------------------------------------------------------------\n"
	s += "-------------- Press q to quit · Enter to confirm · Esc to navigate home ---------------"
	s += "\n----------------------------------------------------------------------------------------\n"

	// Send the UI for rendering
	return tea.NewView(s)
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {

	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
