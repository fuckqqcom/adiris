package routers

import (
	"adiris/middleware"
	"adiris/routers/api/v1/blog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func InitRouter() {
	r := iris.New()
	r.Use(recover.New())
	r.Use(middleware.Check)
	v1 := r.Party("/v1")
	PingParty(v1)

	v2 := r.Party("/v2")
	PongParty(v2)
	r.Run(iris.Addr(":8080"))
}

func PingParty(r iris.Party) {
	r.Get("/ping", blog.Ping)
}

func PongParty(r iris.Party) {
	r.Get("/pong", blog.Pong)
}
