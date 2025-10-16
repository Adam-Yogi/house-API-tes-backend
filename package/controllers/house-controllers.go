package controllers

import (
	"encoding/json"
	"fmt"
	"gin-api/package/models"
	"gin-api/package/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewHouse models.House

func GetHouse(w http.ResponseWriter, r *http.Request) {
	newHouses := models.GetAllHouse()
	res, _ := json.Marshal(newHouses)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHouseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	houseId := vars["houseId"]
	ID, err := strconv.ParseInt(houseId, 0, 0)
	if err != nil {
		fmt.Println("Error when parsing")
	}
	houseDetails, _ := models.GetHouseById(ID)
	res, _ := json.Marshal(houseDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateHouse(w http.ResponseWriter, r *http.Request) {
	CreateHouse := &models.House{}
	utils.ParseBody(r, CreateHouse)
	h := CreateHouse.CreateHouse()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	houseId := vars["houseId"]
	ID, err := strconv.ParseInt(houseId, 0, 0)
	if err != nil {
		fmt.Println("Error when parsing")
	}
	house := models.DeleteHouse(ID)
	res, _ := json.Marshal(house)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateHouse(w http.ResponseWriter, r *http.Request) {
	updateHouse := &models.House{}
	utils.ParseBody(r, updateHouse)
	vars := mux.Vars(r)
	houseId := vars["houseId"]
	ID, err := strconv.ParseInt(houseId, 0, 0)
	if err != nil {
		fmt.Println("Error when parsing")
	}
	houseDetails, db := models.GetHouseById(ID)

	if updateHouse.Title != "" {
		houseDetails.Title = updateHouse.Title
	}
	if updateHouse.Poster != "" {
		houseDetails.Poster = updateHouse.Poster
	}
	if updateHouse.Description != "" {
		houseDetails.Description = updateHouse.Description
	}

	db.Save(&houseDetails)
	res, _ := json.Marshal(houseDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
