// Generated by gopygen; DO NOT EDIT
package game

import (
	"fmt"
	"gem/game/server"
	"gem/runite"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var UpdateServiceDef = py.Class{
	Name:    "UpdateService",
	Pointer: (*UpdateService)(nil),
}

// Registers this type with a python module
func RegisterUpdateService(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = UpdateServiceDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("UpdateService", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func (obj UpdateService) Alloc() (*UpdateService, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := UpdateServiceDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*UpdateService)
	// Copy fields

	alloc.runite = obj.runite

	alloc.queue = obj.queue

	alloc.notify = obj.notify

	return alloc, nil
}

func (svc *UpdateService) Py_Init(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Init: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "*runite.Context")
	if err != nil {
		return nil, err
	}

	svc.Init(in_0.(*runite.Context))

	py.None.Incref()
	return py.None, nil

}

func (svc *UpdateService) Py_NewClient(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 2 {
		return nil, fmt.Errorf("Py_NewClient: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "*server.Connection")
	if err != nil {
		return nil, err
	}

	in_1, err := gopygen.TypeConvIn(args[1], "int")
	if err != nil {
		return nil, err
	}

	res0 := svc.NewClient(in_0.(*server.Connection), in_1.(int))

	out_0, err := gopygen.TypeConvOut(res0, "Client")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (svc *UpdateService) Py_processQueue(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_processQueue: parameter length mismatch")
	}

	svc.processQueue()

	py.None.Incref()
	return py.None, nil

}
