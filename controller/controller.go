package controller

import (
	U "LoginSystem/until"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	
	//for check mail 
	"github.com/badoux/checkmail"

	// For mysql
	_ "github.com/go-sql-driver/mysql"

	//for crpyt
	"golang.org/x/crypto/bcrypt"
)

// Home function
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, username := U.AllSessions(r)
	loggedIn(w, "/welcome", r)
	renderTemplates(w, "index", &Page{"Home", id, username})
}

// Welcome function
func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "welcome", &Page{"Welcome", nil, nil})
}

// Error function
func Error(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplates(w, "error", &Page{"Oops! Something went wrong", nil, nil})
}

// Register function
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "register", &Page{"Register to continue!", nil, nil})
}

// Login function
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "login", &Page{"Login to continue!", nil, nil})
}

// Profile function
func Profile(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	loggedIn(w, "", r)
	id, username := U.AllSessions(r)
	renderTemplates(w, "profile", &Page{"Profile Page", id, username})
}

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

	username := strings.TrimSpace(r.PostFormValue("username"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	passwordAgain := strings.TrimSpace(r.PostFormValue("password_again"))
	//check mail 
	mailErr := checkmail.ValidateFormat(email)

	db := U.DB()

	var (
		userCount  int
		emailCount int
	)

	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&userCount)
	db.QueryRow("SELECT COUNT(id) AS emailCount FROM users WHERE email=?", email).Scan(&emailCount)

	if username == "" || email == "" || password == "" || passwordAgain == "" {
		Response["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		Response["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		Response["mssg"] = "Invalid Format!"
	} else if password != passwordAgain {
		Response["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		Response["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		Response["mssg"] = "Email already exists!"
	} else {

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		_, insErr := db.Exec(
			"INSERT INTO users(username, email, password) VALUES(?, ?, ?) ",
			username,
			email,
			hash,
		)
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

	rusername := strings.TrimSpace(r.PostFormValue("username"))
	rpassword := strings.TrimSpace(r.PostFormValue("password"))

	db := U.DB()
	var (
		userCount int
		id        int
		username  string
		password  string
	)

	db.QueryRow("SELECT COUNT(id) AS userCount, id, username, password FROM users WHERE username=?", rusername).Scan(&userCount, &id, &username, &password)

	if rusername == "" || rpassword == "" {
		Response["mssg"] = "Some values are missing!"
	} else if userCount == 0 {
		Response["mssg"] = "Invalid username!"
	} else if encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword)); encErr != nil {
		Response["mssg"] = "Invalid password!"
	} else {

		session := U.GetSession(r)
		session.Values["id"] = id
		session.Values["username"] = username
		session.Save(r, w)

		Response["success"] = true
		Response["mssg"] = "You are now logged in"
		Response["user"] = id

	}

	final, err := json.Marshal(Response)
	if err != nil {
		panic(err)
	}

	fmt.Println(userCount)
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)

}
