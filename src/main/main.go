package main

import (
	"fmt"
	c "go-userlogin/controller"
	u "go-userlogin/until"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	controller := httprouter.New()
	controller.POST("/user/register", c.UserRegister)
	controller.POST("/user/login", c.UserLogin)

	port := u.GetElement("port")
	fmt.Println("login system start...")
	server := negroni.Classic()
	server.UseHandler(controller)
	server.Run(":" + port)
}
