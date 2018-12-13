package blog

import (
	"adiris/pkg/app"
	"fmt"
)

func Pong(c app.Eng) {
	r := app.R{c}
	fmt.Println("user--->", c.Values().Get("user"))
	fmt.Println("value--->", c.Values().Get("value"))
	fmt.Println("user--->", c.Values().Get("user"))

	data := make(map[string]interface{})
	data["pong"] = "pong"
	r.Response(data)
}

func Ping(c app.Eng) {
	r := app.R{c}
	data := make(map[string]interface{})
	data["ping"] = "ping"
	r.Response(data)
}
