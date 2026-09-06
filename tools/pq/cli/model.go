package cli

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var warningStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("196")).
	Align(lipgloss.Center).
	Padding(1, 10)

var dialogStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("241")).
	Padding(1, 4)

var buttonStyle = lipgloss.NewStyle()

var buttonSelectedStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("57"))

var readOnlyMessageActions = []string{"<- Back", "Show payload"}
var writeMessageActions = []string{"Requeue", "Ack (drop)"}

var dialogActions = []string{"Cancel", "Confirm"}

type MessagesUpdated struct {
	Messages []Message
}

type DialogResult struct {
	Err error
}

func (m Model) FetchMessages() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) WaitForMessages() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	backend Backend
	sub     chan MessagesUpdated

	chosenMessage     *Message
	chosenMessageGone bool

	table    table.Model
	messages []Message

	chosenAction  int
	currentDialog *Dialog

	showingPayload  bool
	payloadViewport viewport.Model
}

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func NewModel(backend Backend) Model { _ = "STUB: not implemented"; return *new(Model) }

type Dialog struct {
	Prompt  string
	Action  func() tea.Msg
	Choice  int
	Running bool
}
