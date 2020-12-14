package solver

import (
	"image"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
)

// A solver is capable of computing the solution of the exit calendar.
type solver struct {
	decoderBoard calendar.DecoderBoard
	calendar     calendar.Calendar
}

// Create a new solver with given decoder board and calendar.
func NewSolver(decoderBoard calendar.DecoderBoard, calendar calendar.Calendar) *solver {
	return &solver{
		decoderBoard: decoderBoard,
		calendar:     calendar,
	}
}

// Solve the calendar up to given day.
// For example, if day is set to 13, then doors 1 to 13 will be solved.
func (s *solver) SolveToDay(day uint8) Path {
	if day > 24 {
		day = 24
	}

	paths := s.traverse(Path{s.calendar.GetStart()}, day)
	return paths
}

// Internally used function to do depth-first traversal.
// If no path is found, it panics because this should never happen.
func (s *solver) traverse(p Path, maxDay uint8) Path {
	if len(p) == int(maxDay) {
		return p
	}

	for x := 0; x < s.calendar.GetWidth(); x++ {
		for y := 0; y < s.calendar.GetHeight(); y++ {
			tail := p.getTail()

			doorCoordinates := image.Pt(x, y)
			iconsOnDoor := s.calendar.GetIconsOnDoor(doorCoordinates)

			directionsOnDoor := s.decoderBoard.GetDirectionByIcons(iconsOnDoor)
			aggregatedDirection := image.Point(directionsOnDoor[0].Add(directionsOnDoor[1]).Add(directionsOnDoor[2]))

			if doorCoordinates == tail.Add(aggregatedDirection) {
				p = p.appendPoint(doorCoordinates)
				return s.traverse(p, maxDay)
			}
		}
	}

	panic("No path found!")
}
