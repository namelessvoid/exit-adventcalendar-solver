package main

import (
	"fmt"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
	"github.com/namelessvoid/exit-solver/internal/pkg/solver"
)

func main() {
	fmt.Println("Running solver for exit advent calendar.")

	decoderBoard := calendar.NewDecoderBoard()
	calendar := calendar.NewCalendar()

	solver := solver.NewSolver(decoderBoard, calendar)
	fmt.Println(solver.Solve())
}
