package main

import (
	"fmt"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
)

func main() {
	fmt.Println("Running solver for exit advent calendar.")

	decoderBoard := calendar.NewDecoderBoard()
	calendar := calendar.NewCalendar()

	solver := newSolver(decoderBoard, calendar)
	fmt.Println(solver.solve())
}
