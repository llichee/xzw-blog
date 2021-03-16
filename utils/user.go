package utils

import "xzw-blog/models"

func GetUserByname(username string) *models.User {
	var u = new(models.User)
	models.DB.Find(&u, "username=?", username)
	if u != nil {
		return u
	}
	return nil
}

func Auth(username, password string) *models.User {
	if user := GetUserByname(username); user == nil {
		return nil
	} else {
		if user.Password == password {
			return user
		} else {
			return nil
		}
	}
}
