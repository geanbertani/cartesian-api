package main

import (
	"github.com/geanbertani/cartesian-api/libs/api"
	"github.com/geanbertani/cartesian-api/services/points"
)

func main() {
	api.Make()
	api.Instance(points.AddRoutes)
	api.Run()
}
