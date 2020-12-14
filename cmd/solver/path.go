package main

import "image"

// path defines the order in which the doors are opened
type path []image.Point

func (path path) getTail() image.Point {
	if len(path) == 0 {
		return image.Point{}
	}

	return path[len(path)-1]
}

func (path path) appendPoint(point image.Point) path {
	return append([]image.Point(path), point)
}
