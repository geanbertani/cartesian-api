package points

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointsDistance struct {
	Point
	Distance int `json:"distance"`
}
