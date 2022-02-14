package points

import (
	"errors"

	"github.com/geanbertani/cartesian-api/libs/env"
	"github.com/geanbertani/cartesian-api/libs/json"
	"github.com/geanbertani/cartesian-api/settings"
	"github.com/geanbertani/cartesian-api/settings/postgres"
	"github.com/jmoiron/sqlx"
)

func GetPointsByDistance(filter Filter) (points []Point, err error) {
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

func getPointsByDistance(point Point, distance int) (points []Point, err error) {
	tx := postgres.MustGetByFile(settings.POSTGRES_FILE).MustBegin()
	defer tx.Rollback()

	points, err = loadPoints()
	if err != nil {
		return
	}

	err = AddPointsTx(points, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func AddPointsTx(points []Point, tx *sqlx.Tx) (err error) {
	return addPointsTx(points, tx)
}

func loadPoints() (points []Point, err error) {
	env.MustSetByJSONFile(settings.ENVIRONMENT_FILE)

	filePoints := env.MustString("points_data_file")

	err = json.UnmarshalFile(filePoints, &points)

	return
}
