package points

import (
	"errors"
	"fmt"
	"os"
)

func GetPointsByDistance(filter Filter) (points []PointsDistance, err error) {
	point := Point{
		X: filter.X,
		Y: filter.Y,
	}

	distance := filter.Distance

	// fmt.Printf("debug: \n point: %v \n distance: %v\n\n", point, distance)

	points, err = getPointsByDistance(point, distance)
	if err != nil {
		return
	}

	if len(points) <= 0 {
		err = errors.New("no points found")
	}

	return
}

func getPointsByDistance(point Point, distance int) (points []PointsDistance, err error) {

	return
}

func loadPoints() (err error) {
	filepath := "data/points.json"

	file, err := os.Open(filepath)
	if err != nil {
		return
	}

	fmt.Printf("debug: \n %v \n\n", file)

	return
}
