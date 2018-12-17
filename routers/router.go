package routers

import (
	"adiris/middleware"
	"adiris/routers/api/admin/v1"
	"adiris/routers/api/v1/blog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func InitRouter() {
	r := iris.New()
	r.Use(recover.New())
	r.Use(middleware.Check)
	a := r.Party("/v1")
	Admin(a)

	//v2 := r.Party("/v2")
	//PongParty(v2)
	r.Run(iris.Addr(":8080"))
}

func Admin(r iris.Party) {
	r.Get("/register", v1.Register)
}

func PongParty(r iris.Party) {
	r.Get("/pong", blog.Pong)
}
