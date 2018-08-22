package controller

import (
	. "../model"
	. "../utils"
	."../service"

	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)


var productModel = ProductModel{}

// GET product
func IndexProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limit, page := 20, 1
	query := bson.M{}
	for k, v := range r.Form {
		switch k {
		case "limit":
			if pageSize, err := strconv.Atoi(v[0]); err == nil {
				limit = pageSize
			}
		case "page":
			if pageNum, err := strconv.Atoi(v[0]); err == nil {
				page = pageNum
			}
		case "barcode":
			query[k] = bson.M{"$regex": bson.RegEx{Pattern: v[0], Options: "i"}}
		default:
			query[k] = v[0]
		}
	}
	products, err := productModel.ProductsFindByParams(query, limit, page)
	total, err := productModel.ProductsNum()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, Msg{"code": 200, "total": total, "data": products})
}

// GET a product by its ID
func ShowProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := productModel.ProductFindById(params["productId"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "查询失败")
		return
	}
	RespondWithJson(w, http.StatusOK, product)
}

// POST a new product
func CreateProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("CofoxAPI: %s", r.Body)
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	product.ID = bson.NewObjectId()
	if err := productModel.ProductInsert(product); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, product)
}

// PUT update an existing product
func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	if err := productModel.ProductUpdate(product); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, Msg{"code": 200, "msg": "更新成功"})
}

// DELETE an existing product
func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		RespondWithError(w, http.StatusBadRequest, "转码失败")
		return
	}
	if err := productModel.ProductDelete(product); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, Msg{"code": 200, "msg": "删除成功"})
}
