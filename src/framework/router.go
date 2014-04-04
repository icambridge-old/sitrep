package framework

import (
	"reflect"
)

type Router struct {
	 Controllers  map[string]reflect.Type
}

func (r *Router) registerController(c interface {}) {

	t := reflect.TypeOf(c)
	r.Controllers[t.Name()] = t

}
