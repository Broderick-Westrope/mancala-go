package tui

import (
	"fmt"
	"slices"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/charmbracelet/lipgloss"
)

const third = 1.0 / 3.0

func (m boardModel) buildBoard(message string) string {
	numberMaxLength := 5
	pitTemplate := ".-=-=-=-.\n|       |\n| %s |\n`-=-=-=-'"
	storeTemplate := ".-=-=-=-.\n|       |\n|       |\n|       |\n| %s |\n|       |\n|       |\n|       |\n`-=-=-=-'"

	topSide := m.game.Side1
	bottomSide := m.game.Side2
	var isTopTurn bool
	if m.game.Turn == mancala.Player1Turn {
		isTopTurn = true
	} else {
		isTopTurn = false
	}

	// TODO: Look into whether this should use panic or not
	if len(topSide.Pits) != len(bottomSide.Pits) {
		panic("top and bottom sides have different number of pits")
	} else if len(topSide.Pits) < 1 {
		panic("sides have no pits")
	}

	// Reverse the top pits
	pits := make([]int, len(topSide.Pits))
	copy(pits, topSide.Pits)
	slices.Reverse(pits)

	// Create the top row of pits
	var renderPit func(...string) string
	if isTopTurn {
		renderPit = baseStyle.Render
	} else {
		renderPit = renderDisabled
	}
	var topPits string
	for i, pit := range pits {
		thisPit := lipgloss.PlaceHorizontal(numberMaxLength, lipgloss.Center, fmt.Sprint(pit))
		thisPit = fmt.Sprintf(pitTemplate, thisPit)
		if isTopTurn && i == m.cursor {
			thisPit = renderSelected(thisPit)
		} else {
			thisPit = renderPit(thisPit)
		}
		topPits = lipgloss.JoinHorizontal(lipgloss.Center, topPits, thisPit)
	}

	// Create the middle text
	middle := lipgloss.PlaceHorizontal(lipgloss.Width(topPits), lipgloss.Center, message)

	// Create the bottom row of pits
	if !isTopTurn {
		renderPit = baseStyle.Render
	} else {
		renderPit = renderDisabled
	}
	var bottomPits string
	for i, pit := range bottomSide.Pits {
		thisPit := lipgloss.PlaceHorizontal(numberMaxLength, lipgloss.Center, fmt.Sprint(pit))
		thisPit = fmt.Sprintf(pitTemplate, thisPit)
		if !isTopTurn && i == m.cursor {
			thisPit = renderSelected(thisPit)
		} else {
			thisPit = renderPit(thisPit)
		}
		bottomPits = lipgloss.JoinHorizontal(lipgloss.Center, bottomPits, thisPit)
	}

	// Join the top and bottom pit rows with the middle text
	board := lipgloss.JoinVertical(lipgloss.Center, topPits, middle, bottomPits)

	// Add the stores and side borders
	leftStore := lipgloss.PlaceHorizontal(numberMaxLength, lipgloss.Center, fmt.Sprint(topSide.Store))
	leftStore = fmt.Sprintf(storeTemplate, leftStore)
	rightStore := lipgloss.PlaceHorizontal(numberMaxLength, lipgloss.Center, fmt.Sprint(bottomSide.Store))
	rightStore = fmt.Sprintf(storeTemplate, rightStore)
	if isTopTurn {
		leftStore = baseStyle.Render(leftStore)
		rightStore = renderDisabled(rightStore)
	} else {
		leftStore = renderDisabled(leftStore)
		rightStore = baseStyle.Render(rightStore)
	}
	renderedSide := baseStyle.Render("!\n¡\n|\n!\n:\n¡\n|\n!\n¡")
	board = lipgloss.JoinHorizontal(lipgloss.Center, renderedSide, leftStore, board, rightStore, renderedSide)

	// Format top name and build top border
	borderWidth := 2
	if m.topBorder == "" {
		m.buildTop(lipgloss.Width(board)-baseStyle.GetHorizontalPadding(), borderWidth, topSide.Player.GetName())
	}
	// Format top name and add top border
	if m.bottomBorder == "" {
		m.buildBottom(lipgloss.Width(board)-baseStyle.GetHorizontalPadding(), borderWidth, bottomSide.Player.GetName())
	}
	// Add top and bottom borders to the board
	board = lipgloss.JoinVertical(lipgloss.Left, m.topBorder, board, m.bottomBorder)

	return board
}

func (m *boardModel) buildTop(width, borderWidth int, name string) {
	var top string
	firstThird := int(float64(width) * third)
	secondThird := width - firstThird

	for i := 0; i < width; i++ {
		switch i {
		case 0, width - 1:
			top += "."
		case firstThird, secondThird:
			top += "<"
		default:
			if i%2 == 0 {
				top += "-"
			} else {
				top += "="
			}
		}
	}

	nameWidth := width - borderWidth
	if len(name) > nameWidth {
		name = name[:nameWidth-3] + "..."
	}

	name = lipgloss.PlaceHorizontal(nameWidth, lipgloss.Center, name)
	name = lipgloss.JoinHorizontal(lipgloss.Center, "|", name, "|")
	m.topBorder = lipgloss.JoinVertical(lipgloss.Left, baseStyle.Render(top), baseStyle.Render(name))
}

func (m *boardModel) buildBottom(width, borderWidth int, name string) {
	nameWidth := width - borderWidth
	if len(name) > nameWidth {
		name = name[:nameWidth-3] + "..."
	}

	name = lipgloss.PlaceHorizontal(nameWidth, lipgloss.Center, name)
	name = lipgloss.JoinHorizontal(lipgloss.Center, "|", name, "|")

	var bottom string
	firstThird := int(float64(width) * third)
	secondThird := width - firstThird

	for i := 0; i < width; i++ {
		switch i {
		case 0:
			bottom += "`"
		case width - 1:
			bottom += "'"
		case firstThird, secondThird:
			bottom += ">"
		default:
			if i%2 == 0 {
				bottom += "-"
			} else {
				bottom += "="
			}
		}
	}

	m.bottomBorder = lipgloss.JoinVertical(lipgloss.Left, baseStyle.Render(name), baseStyle.Render(bottom))
}
