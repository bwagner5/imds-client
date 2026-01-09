/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tui

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/bwagner5/imds/pkg/docs"
	"github.com/bwagner5/imds/pkg/imds"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Padding(0, 1)

	pathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)

	valueBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39")).
			Padding(1, 2)

	valueTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			MarginBottom(1)

	valueContentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("114"))

	fullscreenTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("229")).
				Background(lipgloss.Color("57")).
				Padding(0, 1).
				MarginBottom(1)

	fullscreenValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("114")).
				Padding(1, 2)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)

	searchBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("212")).
			Padding(0, 1)
)

type item struct {
	name     string
	desc     string
	value    string
	isDir    bool
	path     string
	fullPath string // for search results
}

func (i item) Title() string {
	if i.fullPath != "" {
		// Search result - show full path
		if i.isDir {
			return "📁 " + i.fullPath + "/"
		}
		return "   " + i.fullPath
	}
	if i.isDir {
		return "📁 " + i.name + "/"
	}
	return "   " + i.name
}

func (i item) Description() string {
	if i.desc != "" {
		if len(i.desc) > 80 {
			return i.desc[:77] + "..."
		}
		return i.desc
	}
	return ""
}

func (i item) FilterValue() string {
	if i.fullPath != "" {
		return i.fullPath
	}
	return i.name
}

type Model struct {
	client       *imds.Client
	ctx          context.Context
	data         map[string]any
	list         list.Model
	path         []string
	width        int
	height       int
	docs         map[string]string
	showingValue bool
	selectedItem item
	searchMode   bool
	searchInput  textinput.Model
	allItems     []item // flattened list of all items for search
}

func New(ctx context.Context, client *imds.Client) *Model {
	m := &Model{
		client: client,
		ctx:    ctx,
		width:  80,
		height: 24,
		docs:   buildDocsMap(),
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("229")).Bold(true)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("114"))
	delegate.SetHeight(2)

	m.list = list.New([]list.Item{}, delegate, 80, 14)
	m.list.Title = "IMDS Explorer"
	m.list.Styles.Title = titleStyle
	m.list.SetShowStatusBar(true)
	m.list.SetFilteringEnabled(false) // We handle filtering ourselves
	m.list.SetShowHelp(false)

	m.searchInput = textinput.New()
	m.searchInput.Placeholder = "Search all keys..."
	m.searchInput.CharLimit = 100

	return m
}

func buildDocsMap() map[string]string {
	m := make(map[string]string)
	for _, e := range docs.InstanceMetadataCategoryEntries {
		m["meta-data/"+e.Category] = e.Description
	}
	for _, e := range docs.DynamicCategoryEntries {
		m["dynamic/"+e.Category] = e.Description
	}
	return m
}

type dataLoaded struct {
	data map[string]any
}

func (m *Model) Init() tea.Cmd {
	return func() tea.Msg {
		data := m.client.GetAll(m.ctx, "")
		return dataLoaded{data: data}
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		listHeight := msg.Height - 12
		if listHeight < 5 {
			listHeight = 5
		}
		m.list.SetSize(msg.Width, listHeight)

	case dataLoaded:
		m.data = msg.data
		m.buildAllItems()
		m.updateList()

	case tea.KeyMsg:
		// Handle fullscreen value view
		if m.showingValue {
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			default:
				m.showingValue = false
				return m, nil
			}
		}

		// Handle search mode
		if m.searchMode {
			switch msg.String() {
			case "esc":
				m.searchMode = false
				m.searchInput.SetValue("")
				m.updateList()
				return m, nil
			case "enter":
				if sel, ok := m.list.SelectedItem().(item); ok {
					if sel.isDir {
						m.path = strings.Split(sel.fullPath, "/")
						m.searchMode = false
						m.searchInput.SetValue("")
						m.updateList()
						m.list.ResetSelected()
					} else {
						m.selectedItem = sel
						m.showingValue = true
					}
				}
				return m, nil
			case "up", "down":
				var cmd tea.Cmd
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			default:
				var cmd tea.Cmd
				m.searchInput, cmd = m.searchInput.Update(msg)
				m.filterSearch()
				return m, cmd
			}
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "/":
			m.searchMode = true
			m.searchInput.Focus()
			m.filterSearch()
			return m, textinput.Blink
		case "enter", " ", "right", "l":
			if sel, ok := m.list.SelectedItem().(item); ok {
				if sel.isDir {
					m.path = append(m.path, sel.name)
					m.updateList()
					m.list.ResetSelected()
				} else {
					m.selectedItem = sel
					m.showingValue = true
				}
			}
		case "backspace", "left", "h":
			if len(m.path) > 0 {
				m.path = m.path[:len(m.path)-1]
				m.updateList()
				m.list.ResetSelected()
			}
		case "esc":
			if len(m.path) > 0 {
				m.path = nil
				m.updateList()
				m.list.ResetSelected()
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.showingValue {
		return m.renderValueFullscreen()
	}

	var b strings.Builder

	// Search box if in search mode
	if m.searchMode {
		b.WriteString(searchBoxStyle.Width(m.width - 4).Render("🔍 " + m.searchInput.View()))
		b.WriteString("\n")
	}

	// List
	b.WriteString(m.list.View())
	b.WriteString("\n")

	// Path breadcrumb (only when not searching)
	if !m.searchMode {
		pathStr := "/"
		if len(m.path) > 0 {
			pathStr = "/" + strings.Join(m.path, "/")
		}
		b.WriteString(pathStyle.Render("📍 " + pathStr))
		b.WriteString("\n")
	}

	// Value preview panel
	b.WriteString(m.renderValuePreview())

	// Help
	if m.searchMode {
		b.WriteString(helpStyle.Render("↑↓ navigate • enter select • esc cancel search"))
	} else {
		b.WriteString(helpStyle.Render("↑↓ navigate • enter/space select • ←/backspace back • / search • q quit"))
	}

	return b.String()
}

func (m *Model) renderValuePreview() string {
	sel, ok := m.list.SelectedItem().(item)
	if !ok || sel.isDir {
		return valueBoxStyle.Width(m.width-4).Height(3).Render(
			valueTitleStyle.Render("Select a key to view its value"),
		) + "\n"
	}

	val := sel.value
	maxLen := (m.width - 8) * 3
	if len(val) > maxLen {
		val = val[:maxLen-3] + "..."
	}

	content := valueTitleStyle.Render(sel.name) + "\n" + valueContentStyle.Render(val)
	return valueBoxStyle.Width(m.width-4).Render(content) + "\n"
}

func (m *Model) renderValueFullscreen() string {
	var b strings.Builder

	title := fmt.Sprintf(" %s ", m.selectedItem.name)
	b.WriteString(fullscreenTitleStyle.Width(m.width).Render(title))
	b.WriteString("\n")

	if m.selectedItem.desc != "" {
		b.WriteString(pathStyle.Render(m.selectedItem.desc))
		b.WriteString("\n\n")
	}

	val := m.selectedItem.value
	wrapped := wrapText(val, m.width-6)
	b.WriteString(fullscreenValueStyle.Render(wrapped))

	lines := strings.Count(b.String(), "\n")
	for i := lines; i < m.height-2; i++ {
		b.WriteString("\n")
	}

	b.WriteString(helpStyle.Render("Press any key to go back"))

	return b.String()
}

func wrapText(s string, width int) string {
	if width <= 0 {
		return s
	}
	var result strings.Builder
	for _, line := range strings.Split(s, "\n") {
		for len(line) > width {
			result.WriteString(line[:width])
			result.WriteString("\n")
			line = line[width:]
		}
		result.WriteString(line)
		result.WriteString("\n")
	}
	return strings.TrimSuffix(result.String(), "\n")
}

func (m *Model) buildAllItems() {
	m.allItems = nil
	m.collectItems(m.data, "")
}

func (m *Model) collectItems(data map[string]any, prefix string) {
	for k, v := range data {
		fullPath := k
		if prefix != "" {
			fullPath = prefix + "/" + k
		}
		desc := m.docs[fullPath]

		switch val := v.(type) {
		case map[string]any:
			m.allItems = append(m.allItems, item{
				name:     k,
				desc:     desc,
				isDir:    true,
				path:     fullPath,
				fullPath: fullPath,
			})
			m.collectItems(val, fullPath)
		case string:
			m.allItems = append(m.allItems, item{
				name:     k,
				desc:     desc,
				value:    val,
				path:     fullPath,
				fullPath: fullPath,
			})
		case []any:
			m.allItems = append(m.allItems, item{
				name:     k,
				desc:     desc,
				value:    formatArray(val),
				path:     fullPath,
				fullPath: fullPath,
			})
		default:
			m.allItems = append(m.allItems, item{
				name:     k,
				desc:     desc,
				value:    fmt.Sprintf("%v", v),
				path:     fullPath,
				fullPath: fullPath,
			})
		}
	}
}

func (m *Model) filterSearch() {
	query := strings.ToLower(m.searchInput.Value())
	if query == "" {
		// Show all items when search is empty
		items := make([]list.Item, len(m.allItems))
		for i, it := range m.allItems {
			items[i] = it
		}
		m.list.SetItems(items)
		m.list.Title = "All Keys"
		return
	}

	var matches []item
	for _, it := range m.allItems {
		pathLower := strings.ToLower(it.fullPath)
		nameLower := strings.ToLower(it.name)
		if strings.Contains(pathLower, query) || strings.Contains(nameLower, query) {
			matches = append(matches, it)
		}
	}

	// Sort by relevance: exact name match first, then by path length
	sort.Slice(matches, func(i, j int) bool {
		iName := strings.ToLower(matches[i].name)
		jName := strings.ToLower(matches[j].name)
		iExact := iName == query
		jExact := jName == query
		if iExact != jExact {
			return iExact
		}
		iPrefix := strings.HasPrefix(iName, query)
		jPrefix := strings.HasPrefix(jName, query)
		if iPrefix != jPrefix {
			return iPrefix
		}
		return len(matches[i].fullPath) < len(matches[j].fullPath)
	})

	items := make([]list.Item, len(matches))
	for i, it := range matches {
		items[i] = it
	}
	m.list.SetItems(items)
	m.list.Title = fmt.Sprintf("Search: %s (%d)", query, len(matches))
}

func (m *Model) updateList() {
	current := m.data
	for _, p := range m.path {
		if next, ok := current[p].(map[string]any); ok {
			current = next
		} else {
			return
		}
	}

	var items []list.Item
	for k, v := range current {
		fullPath := strings.Join(append(m.path, k), "/")
		desc := m.docs[fullPath]

		switch val := v.(type) {
		case map[string]any:
			items = append(items, item{name: k, desc: desc, isDir: true, path: fullPath})
		case string:
			items = append(items, item{name: k, desc: desc, value: val, path: fullPath})
		case []any:
			items = append(items, item{name: k, desc: desc, value: formatArray(val), path: fullPath})
		default:
			items = append(items, item{name: k, desc: desc, value: fmt.Sprintf("%v", v), path: fullPath})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		ii, jj := items[i].(item), items[j].(item)
		if ii.isDir != jj.isDir {
			return ii.isDir
		}
		return ii.name < jj.name
	})

	m.list.SetItems(items)

	if len(m.path) == 0 {
		m.list.Title = "IMDS Explorer"
	} else {
		m.list.Title = m.path[len(m.path)-1]
	}
}

func formatArray(arr []any) string {
	if len(arr) == 0 {
		return "[]"
	}
	var parts []string
	for _, item := range arr {
		parts = append(parts, fmt.Sprintf("%v", item))
	}
	return strings.Join(parts, "\n")
}

func Run(ctx context.Context, client *imds.Client) error {
	p := tea.NewProgram(New(ctx, client), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
