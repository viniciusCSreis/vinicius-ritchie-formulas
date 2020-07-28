package main

import (
	"fmt"
	"github.com/gookit/color"
	"hello/pkg/calc"
	"os"
)

func main() {
	startWork := os.Getenv("START_WORK")
	startLunch := os.Getenv("START_LUNCH")
	endLunch := os.Getenv("END_LUNCH")
	endWork := os.Getenv("END_WORK")

	wh := calc.WorkHours{
		StartWork:  startWork,
		StartLunch: startLunch,
		EndLunch:   endLunch,
		EndWork:    endWork,
	}

	if endWork != "-" {
		workTime, err := wh.Calc()

		if err != nil {
			fmt.Println(color.FgRed.Render(err.Error()))
			os.Exit(1)
		}

		fmt.Println(color.FgGreen.Render(fmt.Sprintf("You work: %v", workTime)))
	}

	printPerfectTime(wh, "08:44")
	printPerfectTime(wh, "08:00")
}

func printPerfectTime(wh calc.WorkHours, target string) {

	result, err := wh.CalcPerfectEndWork(target)

	if err != nil {
		fmt.Println(color.FgRed.Render(err.Error()))
		os.Exit(1)
	}

	msg := fmt.Sprintf("if you want to work: %v you need to end work : %v ", target, result)
	fmt.Println(color.FgLightBlue.Render(msg))
}
