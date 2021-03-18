package utils

import (
	"xzw-blog/models"
)

func GetUserByname(username string) *models.User {
	var u = new(models.User)
	models.DB.Find(&u, "username=?", username)
	if u != nil {
		return u
	} else {
		return nil
	}
}

func Auth(username, password string) bool {
	if u := GetUserByname(username); u.Username == "" {
		return false
	} else {
		if u.Password == Md5(password) {
			return true
		} else {
			return false
		}
	}
}
