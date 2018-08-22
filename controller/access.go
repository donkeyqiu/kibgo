package controller

import (
	. "../utils"

	"net/http"
	"github.com/gorilla/mux"
)

// POST Login
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload, err := userModel.Login(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, payload["msg"].(string))
		return
	}
	RespondWithJson(w, http.StatusOK, payload)
}

// Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userModel.UserFindById(params["userId"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "查询失败")
		return
	}
	RespondWithJson(w, http.StatusOK, user)
}

// 修改密码
func resetPassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload, err := userModel.ResetPassword(r.FormValue("id"), r.FormValue("newPassword"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, payload["msg"].(string))
		return
	}
	RespondWithJson(w, http.StatusOK, payload)
}

