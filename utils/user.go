package utils

import (
	"fmt"
	"xzw-blog/models"
)

func GetUserByname(username string) *models.User {
	var u = new(models.User)
	models.DB.Find(&u, "username=?", username)
	fmt.Println("+++++++++++")
	fmt.Println(u)
	fmt.Println("+++++++++++")
	if u != nil {
		return u
	} else {
		return nil
	}
}

func Auth(username, password string) bool {
	if u := GetUserByname(username); u.Username == "" {
		fmt.Println("=====")
		fmt.Println(u)
		fmt.Println("=====")
		return false
	} else {
		if u.Password == password {
			return true
		} else {
			return false
		}
	}
}
