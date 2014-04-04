package framework

import (
	"testing"
	"reflect"
)

func TestPullRequestsService_GetAll(t *testing.T) {

	type TestController struct {
		Controller
	}

	r := Router{Controllers: map[string]reflect.Type{}}
	r.registerController(TestController{})

	if _, ok := r.Controllers["TestController"]; ok == false {
		t.Error("Expected TestController to be registered it wasn't")
	}
}
