package model

import (
	 u	"LoginSystem/until"
	 "log"
)

type User struct{
	Id int
	Username string	 
	Email string		
	Password string	
}

func Checkuserbyusername(username string)(int64){
	db := u.DB()
	var count int64
	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&count)
	defer db.Close()
	return count
}

func Checkuserbyemail(email string)(int64){
	db := u.DB()
	var count int64
	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE email=?", email).Scan(&count)
	defer db.Close()
	return count
}

func QueryUserByUserName(username string)(User) {
	db := u.DB()
	result := User{}
	rows, _  :=db.Query("SELECT id, username, password FROM users WHERE username=?", username)
	for rows.Next(){
        if err := rows.Scan(&result.Id, &result.Username, &result.Password); err != nil {
            log.Fatal(err)
        }
        //fmt.Printf("name:%s ,id:is %d\n", name, id)
    }
	defer db.Close()
	return result
}


func InsertUser(username string, email string, password []byte) (int64, error) {
	db := u.DB()
	//result := User{}
	
	stmt, err := db.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
    if err != nil {
        log.Fatal("Cannot prepare DB statement", err)
    }

    res, err := stmt.Exec(username, email, password)
    if err != nil {
        log.Fatal("Cannot run insert statement", err)
    }

    id, _ := res.LastInsertId()
//	result, insErr := db.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?) ", username, email, password)
	defer db.Close()
	return id, err
}

func UpdateUserByUsername(username string, email string, password string)(int64, error) {
	db := u.DB()
	stmt, err := db.Prepare("update user set password=? where username = ?")
	 if err != nil {
        log.Fatal("Cannot prepare DB statement", err)
    }

    res, err := stmt.Exec(password, username)
    if err != nil {
        log.Fatal("Cannot run insert statement", err)
    }
	defer db.Close()		
    id, _ := res.LastInsertId()
	return id,err
}