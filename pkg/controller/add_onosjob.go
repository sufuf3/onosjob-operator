package controller

import (
	"github.com/sufuf3/onosjob-operator/pkg/controller/onosjob"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, onosjob.Add)
}
