package controller

import (
	"github.com/example/cdp/pkg/controller/request"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, request.Add)
}
