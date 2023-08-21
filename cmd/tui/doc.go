/*
Mancala-go is a command-line implementation of the game Mancala.
It supports both local two player and single player game modes, where single player is versing one of the included algorithmic opponents.
The game is played in the terminal as an interactive text UI (TUI).
The TUI is built using the Bubble Tea framework and some Bubbles components both built by Charm.

Usage:

	mancala-go [flags]

Flags:

	-h, -help
		Show help message
	-mode string
	    Game type (local, minimax) (default "local")
	-name1 string
	    Name of player 1 (default "Player 1")
	-name2 string
	    Name of player 2 (default "Player 2")
	-pits int
	    Number of pits per side (default 6)
	-stones int
	    Number of stones per pit (default 4)
*/
package main
