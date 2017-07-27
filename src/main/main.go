package main

import (
	C "loginsystem/controller"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"fmt"
)

func main(){
	controller := httprouter.New()
//
//	controller.GET("/", R.Home)
//	controller.GET("/welcome", R.Welcome)
//	controller.GET("/error", R.Error)
//	controller.GET("/register", R.Register)
//	controller.GET("/login", R.Login)
//	controller.GET("/profile/:user", R.Profile)
//	controller.GET("/profile", R.ProfileParamless)
//	controller.GET("/logout", R.Logout)
//
	controller.POST("/user/register", C.UserRegister)
	controller.POST("/user/login", C.UserLogin)
	fmt.Println("login system start...")
	server := negroni.Classic()
	server.UseHandler(controller)
	server.Run(":8080")
}
