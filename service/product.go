package service

import (
	."../model"

	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	BaseModel
}

const (
	PRODUCT = "product"
)
// Find list of products
func (m *ProductModel) AllProducts() ([]Product, error) {
	var products []Product
	err := db.C(PRODUCT).Find(bson.M{}).All(&products)
	return products, err
}

// Find list of products
func (m *ProductModel) ProductsNum() (int, error) {
	countNum, err := db.C(PRODUCT).Find(bson.M{}).Count()
	return countNum, err
}


// Find list of products by params
func (m *ProductModel) ProductsFindByParams(queryMap bson.M, limit int, page int) ([]Product, error) {
	var products []Product
	err := db.C(PRODUCT).Find(&queryMap).Limit(limit).Skip((page - 1)*limit).Sort("-ID").All(&products)
	return products, err
}

// Find a product by its id
func (m *ProductModel) ProductFindById(id string) (Product, error) {
	var user Product
	err := db.C(PRODUCT).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a product into database
func (m *ProductModel) ProductInsert(user Product) error {
	err := db.C(PRODUCT).Insert(&user)
	return err
}

// Delete an existing product
func (m *ProductModel) ProductDelete(user Product) error {
	err := db.C(PRODUCT).Remove(&user)
	return err
}

// Update an existing product
func (m *ProductModel) ProductUpdate(user Product) error {
	err := db.C(PRODUCT).UpdateId(user.ID, &user)
	return err
}
