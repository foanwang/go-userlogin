package model_test

import (
	. "go-userlogin/model"
	"testing"
)

func TestCheckuserbyusername(t *testing.T) {
	var list []string
	list = []string{"foanwang", "foanwang1212", "annie", "annie1219"}

	for _, username := range list {
		userCount := Checkuserbyusername(username)
		if userCount < 0 {
			t.Error("username:%v usercount:%i", username, userCount)
		}
	}
}

func TestCheckuserbyemail(t *testing.T) {
	var list []string
	list = []string{"foanwang@mail.com", "foanwang1212@gmail.com", "annie", "annie1219"}

	for _, email := range list {
		userCount := Checkuserbyemail(email)
		if userCount < 0 {
			t.Error("email:%v usercount:%i", email, userCount)
		}
	}
}

// func TestInsertUser(t *testing.T) {
// 	var list []interface
// 	list = []interface{"foanwang@mail.com", "foanwang1212@gmail.com", "annie","annie1219"}

// 	for _, email := range list{
// 		userCount:=InsertUser(email)
// 		if userCount < 0{
// 			t.Error("email:%v usercount:%i", email, userCount)
// 		}
// 	}
// }
