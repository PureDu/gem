package gem

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterEngine,
}

func init() {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = python.InitModule("gem.engine", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}
}