package models

import (
	"gin-api/package/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type House struct {
	gorm.Model
	Title       string `gorm:""json:"title"`
	Poster      string `json:"poster"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&House{})
}

func (h *House) CreateHouse() *House {
	db.Create(&h)
	return h
}

func GetAllHouse() []House {
	var Houses []House
	db.Find(&Houses)
	return Houses
}

func GetHouseById(Id int64) (*House, *gorm.DB) {
	var getHouse House
	db := db.Where("ID=?", Id).Find(&getHouse)
	return &getHouse, db
}

func DeleteHouse(ID int64) House {
	var house House
	db.Where("ID=?", ID).Delete(&house)
	return house
}
