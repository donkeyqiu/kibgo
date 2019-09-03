package controller

import (
	. "../utils"
	"github.com/gorilla/mux"
	"net/http"
)

// POST Login
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload, err := userModel.Login(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		RespondWithError(w, http.StatusOK, payload["msg"].(string))
		return
	}
	RespondWithJson(w, http.StatusOK, payload)
}

// Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userModel.UserFindById(params["userId"])
	if err != nil {
		RespondWithError(w, http.StatusOK, "查询失败")
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

// 修改密码
func resetPassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload, err := userModel.ResetPassword(r.FormValue("id"), r.FormValue("newPassword"))
	if err != nil {
		RespondWithError(w, http.StatusOK, payload["msg"].(string)) // http.StatusBadRequest
		return
	}
	RespondWithJson(w, http.StatusOK, payload)
}

