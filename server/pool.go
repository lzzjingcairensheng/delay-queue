package sever

import (
	"github.com/Jeffail/tunny"
	"runtime"
)

//Pool routine pool
var Pool *tunny.Pool

func init() {
	numCPUs := runtime.NumCPU()
	Pool = tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		//
		return nil
	})
}
