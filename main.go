package main

import (
	"encoding/json"
	"fmt"
	_ "inventory-api/docs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Item represents the model for an inventory's item
type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

// GetInventory godoc
// @Summary Get all the inventory's items
// @Description Get all the inventory's items
// @Produce json
// @Success 200 {array} Item
// @Router /inventory [get]
func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

// CreateItem godoc
// @Summary Add a new item to the inventory
// @Description Add a new item to the inventory
// @Accept json
// @Produce json
// @Success 200 {array} Item
// @Router /inventory [post]
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)
}

// DeleteItem godoc
// @Summary Delete an existing item from the inventory
// @Description Delete an existing item from the inventory
// @Produce json
// @Success 200 {array} Item
// @Router /inventory/{uid} [delete]
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUID(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

// UpdateItem godoc
// @Summary Update an existing item from the inventory
// @Description Update an existing item from the inventory
// @Produce json
// @Success 200 {array} Item
// @Router /inventory/{uid} [put]
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	//	Delete the item at UID
	_deleteItemAtUID(params["uid"])

	//	Create it with new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUID(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			//	Delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// @title Inventory API
// @version 1.0
// @description Thi is a simple service fon an inventory system
// @host localhost:8000
// @BasePath /
func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese",
		Price: 4.99,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Milk",
		Desc:  "A jug of milk",
		Price: 3.25,
	})
	handleRequests()
}
