package points

type Filter struct {
	X        int `query:"x"`
	Y        int `query:"y"`
	Distance int `query:"distance"`
}
