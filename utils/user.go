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

func Auth(username, password string) (int, *models.User) {
	u := []models.User{}
	models.DB.Where("username=? AND password=?", username, Md5(password)).Find(&u)
	return len(u), &u[0]
}
