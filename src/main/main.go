package main

import (
	C "loginsystem/controller"
	u	"LoginSystem/until"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"fmt"
	 	
)

func main(){
	controller := httprouter.New()
	controller.POST("/user/register", C.UserRegister)
	controller.POST("/user/login", C.UserLogin)
	
	port := u.GetElement("port");
	fmt.Println("login system start...")
	server := negroni.Classic()
	server.UseHandler(controller)
	server.Run(":"+port)
}
