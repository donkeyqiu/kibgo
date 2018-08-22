package service

import (
	."../model"

	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	BaseModel
}

const (
	USER = "user"
)
// Find list of users
func (m *UserModel) AllUsers() ([]User, error) {
	var users []User
	err := db.C(USER).Find(bson.M{}).All(&users)
	return users, err
}

// Find list of users
func (m *UserModel) UsersNum() (int, error) {
	countNum, err := db.C(USER).Find(bson.M{}).Count()
	return countNum, err
}


// Find list of users by params
func (m *UserModel) UsersFindByParams(queryMap bson.M, limit int, page int) ([]User, error) {
	var users []User
	err := db.C(USER).Find(&queryMap).Limit(limit).Skip((page - 1)*limit).Sort("-ID").All(&users)
	return users, err
}

// Find a user by its id
func (m *UserModel) UserFindById(id string) (User, error) {
	var user User
	err := db.C(USER).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Find a user by its username
func (m *UserModel) UserFindByName(username string) (User, error) {
	var user User
	err := db.C(USER).Find(bson.M{"username": username}).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserModel) UserInsert(user User) error {
	err := db.C(USER).Insert(&user)
	return err
}

// Delete an existing user
func (m *UserModel) UserDelete(user User) error {
	err := db.C(USER).Remove(&user)
	return err
}

// Delete many users
func (m *UserModel) UserDeleteAll(selector bson.M) error {
	_, err := db.C(USER).RemoveAll(selector)
	return err
}

// Update an existing user
func (m *UserModel) UserUpdate(user User) error {
	err := db.C(USER).UpdateId(user.ID, &user)
	return err
}

// Update users by selector
func (m *UserModel) UserUpdateAny(selector, updater bson.M) error {
	err := db.C(USER).Update(selector, updater)
	return err
}
