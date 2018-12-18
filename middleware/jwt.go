package middleware

import (
	"adiris/pkg/app"
	"fmt"
)

func Check(c app.Eng) {
	header := c.Request().Header
	token := header.Get("token")
	/**
	参数校验:
		加密
		token解析
		获取权限
	*/

	fmt.Println("token-->", token)
	//m := make(chan map[string]interface{})
	//fmt.Println(m, c)
	//tools.ParseToken(token, m)
	//rm := <-m
	//
	//if rm["error"] == nil{
	//	claims := rm["token"]
	//	fmt.Println(claims)
	//}
	//c.Values().Set("user", "1111")
	//c.Values().Set("value", 100)
	c.Next() // execute the "after" handler registered via `DoneGlobal`.
}
