package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
	"github.com/namelessvoid/exit-solver/internal/pkg/solver"
)

func main() {
	fmt.Println("Running solver for exit advent calendar.")

	dayFlag := flag.Uint("day", 1, "the day up to which the calender should be solved. Must be between 1 and 24")
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 24 {
		flag.Usage()
		os.Exit(1)
	}

	decoderBoard := calendar.NewDecoderBoard()
	calendar := calendar.NewCalendar()

	solver := solver.NewSolver(decoderBoard, calendar)
	path := solver.SolveToDay(uint8(*dayFlag))

	renderPath(calendar, path)
}

func renderPath(c calendar.Calendar, p solver.Path) {
	for y := 0; y < c.GetHeight(); y++ {
		for x := 0; x < c.GetWidth(); x++ {
			day := p.GetDayForDoor(image.Pt(x, y))

			if day == -1 {
				fmt.Print(" ?? ")
			} else {
				fmt.Printf(" %2d ", day)
			}
		}
		fmt.Println()
	}
}
