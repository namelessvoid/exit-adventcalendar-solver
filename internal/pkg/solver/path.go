package solver

import "image"

// Path represents a list of doors in form of their coordinates on the calendar.
type Path []image.Point

func (path Path) getTail() image.Point {
	if len(path) == 0 {
		return image.Point{}
	}

	return path[len(path)-1]
}

func (path Path) appendPoint(point image.Point) Path {
	return append([]image.Point(path), point)
}

// GetDayForDoor returns the position of the given point in the path.
// This corresponds to the day at which the given door will be opened.
func (path Path) GetDayForDoor(point image.Point) int {
	for i, p := range path {
		if p == point {
			return i + 1
		}
	}

	return -1
}
