package controller

import (
	."../model"
	."../utils"
	."../service"

	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var userModel = UserModel{}

// GET users
func IndexUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limit, page:= 20, 1
	query := bson.M{}
	for k, v:= range r.Form {
		switch k {
			case "limit":
				if pageSize, err := strconv.Atoi(v[0]); err == nil {
					limit = pageSize
				}
			case "page":
				if pageNum, err := strconv.Atoi(v[0]); err == nil {
					page = pageNum
				}
			default:
				query[k] = bson.M{"$regex": bson.RegEx{Pattern: v[0], Options: "i"}}
		}
	}
	users, err := userModel.UsersFindByParams(query, limit, page )
	total, err := userModel.UsersNum()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, Msg{"code": 200, "total": total, "data": users})
}

// GET a user by its ID
func ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userModel.UserFindById(params["userId"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "查询失败")
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

// POST a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("CofoxAPI: %s", r.Body)
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	user.ID = bson.NewObjectId()
	if err := userModel.UserInsert(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, user)
}

// PUT update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	if err := userModel.UserUpdate(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK,Msg{"code": 200, "msg": "更新成功"})
}

// DELETE an existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	if err := userModel.UserDelete(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, Msg{"code": 200, "msg": "删除成功"})
}