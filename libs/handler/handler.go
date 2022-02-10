// It's a lib for some utility functions to help with handler in Echo framework
package handler

import (
	"reflect"

	"github.com/eucatur/go-toolbox/validator"
	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

const PARAMETERS = "parameters"
const MESSAGE = "message"

type Handler struct {
	Message string `json:"message" form:"message" query:"message"`
}

// BindAndValidate like the name Validade and bind one struct with the validador golang lib
func BindAndValidate(c echo.Context, obj interface{}, args ...interface{}) (err error) {
	obj = reflect.ValueOf(obj).Elem().Interface()
	obj = reflect.New(reflect.TypeOf(obj)).Interface()

	err = c.Bind(obj)
	if err != nil {
		return
	}

	var options []string
	if len(args) > 0 {
		group, ok := args[0].(string)
		if ok {
			options = append(options, group)
		}
	}

	vErr := validator.Validate(obj, options...)
	if vErr != nil {
		err = c.JSON(422, vErr)
		if err != nil {
			return
		}
		return vErr
	}

	defaults.SetDefaults(obj)

	c.Set(PARAMETERS, obj)

	return
}
