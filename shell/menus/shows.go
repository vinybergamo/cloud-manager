package sheelMenus

import (
	terminal "github.com/nsf/termbox-go"
)

func ShowMenuVertical(options []string) int {
	if err := terminal.Init(); err != nil {
		panic(err)
	}
	defer terminal.Close()

	cursorY := 0

	for {
		terminal.Clear(terminal.ColorDefault, terminal.ColorDefault)

		for i, option := range options {
			fg := terminal.ColorDefault
			bg := terminal.ColorDefault
			if i == cursorY {
				fg = terminal.ColorWhite
				bg = terminal.ColorWhite
			}
			for j, ch := range option {
				terminal.SetCell(j, i, ch, fg, bg)
			}
			terminal.Flush()
		}

		switch ev := terminal.PollEvent(); ev.Type {
		case terminal.EventKey:
			switch ev.Key {
			case terminal.KeyArrowUp:
				if cursorY > 0 {
					cursorY--
				}
			case terminal.KeyArrowDown:
				if cursorY < len(options)-1 {
					cursorY++
				}
			case terminal.KeyEnter:
				return cursorY

			}
		case terminal.EventError:
			panic(ev.Err)
		}
	}
}

func ShowMenuHorizontal(options []string) int {
	if err := terminal.Init(); err != nil {
		panic(err)
	}
	defer terminal.Close()

	cursorX := 0
	cursorY := 0

	for {
		terminal.Clear(terminal.ColorDefault, terminal.ColorDefault)
		for i, option := range options {
			fg := terminal.ColorDefault
			bg := terminal.ColorDefault
			if i == cursorY {
				fg = terminal.ColorWhite
				bg = terminal.ColorBlack
			}
			for j, ch := range option {
				if j == cursorX {
					terminal.SetCell(j, i, ch, fg|terminal.AttrBold, bg)
				} else {
					terminal.SetCell(j, i, ch, fg, bg)
				}
			}
			terminal.Flush()
		}

		switch ev := terminal.PollEvent(); ev.Type {
		case terminal.EventKey:
			switch ev.Key {
			case terminal.KeyArrowUp:
				if cursorY > 0 {
					cursorY--
				}
			case terminal.KeyArrowDown:
				if cursorY < len(options)-1 {
					cursorY++
				}
			case terminal.KeyArrowLeft:
				if cursorX > 0 {
					cursorX--
				}
			case terminal.KeyArrowRight:
				if cursorX < len(options[cursorY])-1 {
					cursorX++
				}
			case terminal.KeyEnter:
				return cursorY
			}
		case terminal.EventError:
			panic(ev.Err)
		}
	}
}
