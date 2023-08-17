Output ./docs/gifs/demo.gif

Set Margin 30
Set MarginFill "#6B50FF"
Set BorderRadius 20
Set Width 1600
Set Height 700
Set FontFamily "FiraCode Nerd Font Mono"

Hide
Type "go build -o tui ./cmd/tui/"
Enter
Type "clear"
Enter
Show

Type "./tui"
Enter
Sleep 3

Type "?"
Sleep 2

Right@500ms 3
Sleep 1
Enter
Sleep 2

Left@500ms 3
Sleep 1
Enter
Sleep 2

Right@500ms 1
Sleep 1
Enter
Sleep 2

Right@500ms 3
Sleep 1
Enter
Sleep 2

Type "q"

Type "./tui -pits 3 -name1 Fry -name2 Bender"
Enter
Sleep 5

Hide
Type "q"
Type "rm tui"
Enter