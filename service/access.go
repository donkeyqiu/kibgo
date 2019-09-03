package service

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

const (
	SecretKey = "Welcome to PurePearl"
)

var userModel = UserModel{}

// Login
func (m *UserModel) Login(username, password string) (Msg, error) {
	user, err := userModel.UserFindByName(username)
	if err != nil {
		return Msg{"code": 0, "msg": "用户名不存在"}, err
	}
	h := md5.New()
	h.Write([]byte(password))
	md5str2 := hex.EncodeToString(h.Sum(nil))
	if user.Password != md5str2[8:24] {
		return Msg{"code": 0, "msg": "密码错误"}, err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*180)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return Msg{"code": 0, "msg": "token生成失败"}, err
	}
	name := "user"
	if user.Username == "admin" {
		name = "admin"
	}
	return Msg{
		"code":             200,
		"status":           "ok",
		"type":             "account",
		"currentAuthority": name,
		"token":            tokenString,
		"userInfo":         user}, err
}

// Logout
func (m *UserModel) Logout() (int, error) {
	countNum, err := db.C(USER).Find(bson.M{}).Count()
	return countNum, err
}

// 修改密码
func (m *UserModel) ResetPassword(id, newPassword string) (Msg, error) {
	_, err := userModel.UserFindById(id)
	if err != nil {
		return Msg{"code": 0, "msg": "用户不存在"}, err
	}
	h := md5.New()
	h.Write([]byte(newPassword))
	md5str2 := hex.EncodeToString(h.Sum(nil))
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	data := bson.M{"$set": bson.M{"password": md5str2[8:24]}}
	if err := userModel.UserUpdateAny(selector, data); err != nil {
		return Msg{"code": 0, "msg": "修改失败"}, err
	}
	return Msg{"code": 200, "msg": "修改成功"}, err
}
