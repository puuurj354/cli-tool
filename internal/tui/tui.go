package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/purnama/scaffold/internal/templates"
)

// ProjectConfig holds the user's project configuration choices
type ProjectConfig struct {
	ProjectName   string
	TemplateName  string
	License       string
	IncludeDocker bool
	InitGit       bool
}

type step int

const (
	stepProjectName step = iota
	stepTemplate
	stepDocker
	stepLicense
	stepGit
	stepDone
)

type model struct {
	step      step
	textInput textinput.Model
	cursor    int
	templates []templates.Template
	licenses  []string
	config    ProjectConfig
	err       error
	width     int
	height    int
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED")).
			MarginBottom(1)

	questionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7C3AED")).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true)
)

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "my-awesome-project"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 40

	return model{
		step:      stepProjectName,
		textInput: ti,
		templates: templates.GetAllTemplates(),
		licenses:  []string{"MIT", "Apache 2.0", "GPL 3.0", "None"},
		config:    ProjectConfig{},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.err = fmt.Errorf("cancelled")
			return m, tea.Quit

		case "enter":
			return m.handleEnter()

		case "up", "k":
			if m.step != stepProjectName && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.step == stepTemplate && m.cursor < len(m.templates)-1 {
				m.cursor++
			} else if m.step == stepLicense && m.cursor < len(m.licenses)-1 {
				m.cursor++
			}

		case "y", "Y":
			if m.step == stepDocker || m.step == stepGit {
				return m.handleYesNo(true)
			}

		case "n", "N":
			if m.step == stepDocker || m.step == stepGit {
				return m.handleYesNo(false)
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	if m.step == stepProjectName {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.step {
	case stepProjectName:
		name := strings.TrimSpace(m.textInput.Value())
		if name == "" {
			name = "my-project"
		}
		m.config.ProjectName = name
		m.step = stepTemplate
		m.cursor = 0

	case stepTemplate:
		m.config.TemplateName = m.templates[m.cursor].Name
		m.step = stepDocker
		m.cursor = 0

	case stepLicense:
		m.config.License = m.licenses[m.cursor]
		m.step = stepGit
		m.cursor = 0
	}

	return m, nil
}

func (m model) handleYesNo(yes bool) (tea.Model, tea.Cmd) {
	switch m.step {
	case stepDocker:
		m.config.IncludeDocker = yes
		m.step = stepLicense
		m.cursor = 0

	case stepGit:
		m.config.InitGit = yes
		m.step = stepDone
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	s.WriteString(titleStyle.Render("🚀 Project Scaffold"))
	s.WriteString("\n\n")

	switch m.step {
	case stepProjectName:
		s.WriteString(questionStyle.Render("? What is the project name?"))
		s.WriteString("\n")
		s.WriteString(m.textInput.View())

	case stepTemplate:
		s.WriteString(questionStyle.Render("? Select a template:"))
		s.WriteString("\n")
		for i, t := range m.templates {
			cursor := "  "
			style := normalStyle
			if i == m.cursor {
				cursor = "▸ "
				style = selectedStyle
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, style.Render(fmt.Sprintf("%s - %s", t.Name, t.Description))))
		}

	case stepDocker:
		s.WriteString(questionStyle.Render("? Include Dockerfile? (y/n)"))

	case stepLicense:
		s.WriteString(questionStyle.Render("? Which license would you like to use?"))
		s.WriteString("\n")
		for i, l := range m.licenses {
			cursor := "  "
			style := normalStyle
			if i == m.cursor {
				cursor = "▸ "
				style = selectedStyle
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, style.Render(l)))
		}

	case stepGit:
		s.WriteString(questionStyle.Render("? Initialize git repository? (y/n)"))

	case stepDone:
		s.WriteString(successStyle.Render("✓ Configuration complete!"))
	}

	s.WriteString("\n\n")
	s.WriteString(normalStyle.Render("(Press Ctrl+C to cancel)"))

	return s.String()
}

// Run starts the interactive TUI and returns the user's configuration
func Run() (ProjectConfig, error) {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		return ProjectConfig{}, err
	}

	finalModel := m.(model)
	if finalModel.err != nil {
		return ProjectConfig{}, finalModel.err
	}

	return finalModel.config, nil
}
