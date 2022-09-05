package db

import (
	"github.com/asdine/storm/v3"
)

var db *storm.DB

type User struct {
	ID       int    `storm:"id,increment"`
	Username string `storm:"unique"`
	Password string
	Groups   []string
}

type Group struct {
	ID            int `storm:"id,increment"`
	Name          string
	AllowedRoutes []string
}

func Setup() error {
	dbTmp, err := storm.Open("/data/credentials.db")
	db = dbTmp
	return err
}

func Close() {
	db.Close()
}

func CreateUser(user User) error {
	return db.Save(&user)
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.One("Username", username, &user)
	return user, err
}

func GetUserByID(id int) (User, error) {
	var user User
	err := db.One("ID", id, &user)
	return user, err
}

func GetUsers() ([]User, error) {
	var users []User
	err := db.All(&users)
	return users, err
}

func DeleteUser(id int) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	return db.DeleteStruct(&user)
}

func UpdateUser(user User) error {
	return db.Update(&user)
}

func CreateGroup(group Group) error {
	return db.Save(&group)
}

func GetGroupByID(id int) (Group, error) {
	var group Group
	err := db.One("ID", id, &group)
	return group, err
}

func GetGroups() ([]Group, error) {
	var groups []Group
	err := db.All(&groups)
	return groups, err
}

func DeleteGroup(id int) error {
	group, err := GetGroupByID(id)
	if err != nil {
		return err
	}
	return db.DeleteStruct(&group)
}

func UpdateGroup(group Group) error {
	return db.Update(&group)
}
