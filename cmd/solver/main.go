package main

import (
	"fmt"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
)

func main() {
	fmt.Println("Running solver for exit advent calendar.")

	decoder := calendar.NewDecoderBoard()
	fmt.Println(decoder)
}
