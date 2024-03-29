package controller

import (
	"encoding/json"
	M "go-userlogin/model"
	U "go-userlogin/until"
	"log"
	"net/http"
	"github.com/badoux/checkmail" ////for check mail
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt" //for crpyt
)

// ProfileParamless function
func ProfileParamless(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)
	http.Redirect(w, r, "/", http.StatusFound)
}

// Logout function
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)
	session := U.GetSession(r)
	delete(session.Values, "id")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}


// UserRegister function
func UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})
	data := M.User{}; 
    if r.Body == nil {
       Response["mssg"] = "Some values are missing!"
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
	  Response["mssg"] = "Some values are missing!"
    }
    
	username := data.Username
	email := data.Email
	password := data.Password
	
	mailErr := checkmail.ValidateFormat(email)
	
	userCount := M.Checkuserbyusername(username)
	emailCount:= M.Checkuserbyemail(email)

	if username == "" || email == "" || password == "" {
		Response["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		Response["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		Response["mssg"] = "Invalid Format!"
	} else if userCount > 0 {
		Response["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		Response["mssg"] = "Email already exists!"
	} else {

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		//insert
		_, insErr:=M.InsertUser(username, email, hash)
		if insErr != nil {
			log.Fatal(insErr)
		}

		Response["mssg"] = "You are now registered!"
		Response["success"] = true

	}

	final, err := json.Marshal(Response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(final)
}

// UserLogin function
func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})
	data := M.User{}; 
    if r.Body == nil {
       Response["mssg"] = "Some values are missing!"
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
	  Response["mssg"] = "Some values are missing!"
    }
	rusername := data.Username
	rpassword := data.Password

	user:= M.QueryUserByUserName(rusername)
	if user.Username == ""{
		Response["mssg"] = "Invalid username!"
	}else if encErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rpassword)); encErr != nil {
		Response["mssg"] = "Invalid password!"
	} else {

		session := U.GetSession(r)
		session.Values["id"] = user.Id
		session.Values["username"] = user.Username
		session.Save(r, w)

		Response["success"] = true
		Response["mssg"] = "You are now logged in"
		Response["user"] = user.Id

	}

	final, err := json.Marshal(Response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)

}
