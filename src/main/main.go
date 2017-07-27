package main

import (
	C "loginsystem/controller"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"fmt"
)

func main(){
	controller := httprouter.New()
	controller.POST("/user/register", C.UserRegister)
	controller.POST("/user/login", C.UserLogin)
	
	fmt.Println("login system start...")
	server := negroni.Classic()
	server.UseHandler(controller)
	server.Run(":8080")
}
