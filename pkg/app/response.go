package app

import "github.com/kataras/iris"

type Eng = iris.Context

type R struct {
	IR iris.Context
}

func (r R) Response(data interface{}) {
	r.IR.JSON(data)
}
