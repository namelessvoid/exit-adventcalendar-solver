package main

import (
	"fmt"
	"image"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
	"github.com/namelessvoid/exit-solver/internal/pkg/solver"
)

func main() {
	fmt.Println("Running solver for exit advent calendar.")

	decoderBoard := calendar.NewDecoderBoard()
	calendar := calendar.NewCalendar()

	solver := solver.NewSolver(decoderBoard, calendar)
	path := solver.SolveToDay(13)

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
