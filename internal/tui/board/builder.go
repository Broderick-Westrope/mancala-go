package board

import (
	"fmt"
	"slices"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/charmbracelet/lipgloss"
)

var (
	storeTemplate   = ".-=-=-=-.\n|       |\n|       |\n|       |\n| %s |\n|       |\n|       |\n|       |\n`-=-=-=-'"
	sideBorder      = "!\n¡\n|\n!\n:\n¡\n|\n!\n¡"
	pitTemplate     = ".-=-=-=-.\n|       |\n| %s |\n`-=-=-=-'"
	topBorder       = ".=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-<-=-=-=-=-=-=-=-=-=-=-=-=-=-=-<-=-=-=-=-=-=-=-=-=-=-=-=-=-=-.\n|   %s                                                                   |"
	bottomBorder    = "|                                                                   %s   |\n`=-=-=-=-=-=-=-=-=-=-=-=-=-=-=->-=-=-=-=-=-=-=-=-=-=-=-=-=-=->-=-=-=-=-=-=-=-=-=-=-=-=-=-bw'\n"
	nameMaxLength   = 20
	numberMaxLength = 5
)

func (m Model) buildBoard() string {
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
	middle := lipgloss.PlaceHorizontal(lipgloss.Width(topPits), lipgloss.Center, m.stopwatch.View())

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
	board = lipgloss.JoinHorizontal(lipgloss.Center, baseStyle.Render(sideBorder), leftStore, board, rightStore, baseStyle.Render(sideBorder))

	// Format top name and add top border
	name := topSide.Player.Name
	if len(name) > nameMaxLength {
		name = name[:nameMaxLength-3] + "..."
	}
	name = lipgloss.PlaceHorizontal(nameMaxLength, lipgloss.Left, name)
	top := fmt.Sprintf(topBorder, name)
	top = baseStyle.Render(top)
	// Format top name and add top border
	name = bottomSide.Player.Name
	if len(name) > nameMaxLength {
		name = name[:nameMaxLength-3] + "..."
	}
	name = lipgloss.PlaceHorizontal(nameMaxLength, lipgloss.Right, name)
	bottom := fmt.Sprintf(bottomBorder, name)
	bottom = baseStyle.Render(bottom)
	// Add top and bottom borders to the board
	board = lipgloss.JoinVertical(lipgloss.Left, top, board, bottom)

	return board
}
