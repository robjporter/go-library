package main

import (
	"fmt"
	. "github.com/robjporter/go-library/aurora"
)

var (
	au = NewAurora(true)
)

func main() {
	fmt.Println("STARTING")
	fmt.Println("SIMPLE===============================================")
	fmt.Println(au.Green("Hello"))
	fmt.Println("SIMPLE EFFECTS=======================================")
	fmt.Println(au.Bold("BOLD"))
	fmt.Println(Italic("Italic"))
	fmt.Println(Underline("Underline"))
	fmt.Println(au.Faint("Faint"))
	fmt.Println("UNCOMMON EFFECTS=====================================")
	fmt.Println(au.DoublyUnderline("DoublyUnderline"))
	fmt.Println(Fraktur("Fraktur"))
	fmt.Println(SlowBlink("SlowBlink"))
	fmt.Println(RapidBlink("RapidBlink"))
	fmt.Println(Blink("Blink"))
	fmt.Println("TEXT EFFECTS=========================================")
	fmt.Println(Reverse("Reverse"))
	fmt.Println(Inverse("Inverse"))
	fmt.Println(Conceal("Conceal"))
	fmt.Println(Hidden("Hidden"))
	fmt.Println("TEXT EFFECTS=========================================")
	fmt.Println(CrossedOut("CrossedOut"))
	fmt.Println(StrikeThrough("StrikeThrough"))
	fmt.Println(Framed("Framed"))
	fmt.Println(Encircled("Encircled"))
	fmt.Println(Overlined("Overlined"))
	fmt.Println("COLOR EFFECTS========================================")
	fmt.Println("Hello, ", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))
	fmt.Printf("Got it %d times\n", Green(1240))
	fmt.Printf("PI is %+1.2e\n", Cyan(3.14))
	fmt.Println(Sprintf(Magenta("Got it %d times"), Green(1240)))
	x := BgMagenta(Bold(Red("x")))
	fmt.Println(x)
	x = Red("x").Bold().BgMagenta()
	fmt.Println(x)
	fmt.Println("COLORS===============================================")
	fmt.Println(Black("Black"))
	fmt.Println(Red("Red"))
	fmt.Println(Green("Green"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(Brown("Brown"))
	fmt.Println(Blue("Blue"))
	fmt.Println(Magenta("Magenta"))
	fmt.Println(Cyan("Cyan"))
	fmt.Println(White("White"))
	fmt.Println(BrightBlack("BrightBlack"))
	fmt.Println(BrightRed("BrightRed"))
	fmt.Println(BrightGreen("BrightGreen"))
	fmt.Println(BrightYellow("BrightYellow"))
	fmt.Println(BrightBlue("BrightBlue"))
	fmt.Println(BrightMagenta("BrightMagenta"))
	fmt.Println(BrightCyan("BrightCyan"))
	fmt.Println(BrightWhite("BrightWhite"))
	fmt.Println("BACKGROUND COLORS====================================")
	fmt.Println(BgBlack("BgBlack"))
	fmt.Println(BgRed("BgRed"))
	fmt.Println(BgGreen("BgGreen"))
	fmt.Println(BgYellow("BgYellow"))
	fmt.Println(BgBrown("BgBrown"))
	fmt.Println(BgBlue("BgBlue"))
	fmt.Println(BgMagenta("BgMagenta"))
	fmt.Println(BgCyan("BgCyan"))
	fmt.Println(BgWhite("BgWhite"))
	fmt.Println(BgBrightBlack("BgBrightBlack"))
	fmt.Println(BgBrightRed("BgBrightRed"))
	fmt.Println(BgBrightGreen("BgBrightGreen"))
	fmt.Println(BgBrightYellow("BgBrightYellow"))
	fmt.Println(BgBrightBlue("BgBrightBlue"))
	fmt.Println(BgBrightMagenta("BgBrightMagenta"))
	fmt.Println(BgBrightCyan("BgBrightCyan"))
	fmt.Println(BgBrightWhite("BgBrightWhite"))
	fmt.Println("SHADE EFFECTS========================================")
	fmt.Println("  ",
		Gray(1-1, " 00-23 ").BgGray(24-1),
		Gray(4-1, " 03-19 ").BgGray(20-1),
		Gray(8-1, " 07-15 ").BgGray(16-1),
		Gray(12-1, " 11-11 ").BgGray(12-1),
		Gray(16-1, " 15-07 ").BgGray(8-1),
		Gray(20-1, " 19-03 ").BgGray(4-1),
		Gray(24-1, " 23-00 ").BgGray(1-1),
	)
}