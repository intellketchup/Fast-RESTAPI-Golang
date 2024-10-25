// handlers/handlers.go
package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "RESTAPI/models"
    "strconv"
)

var items = []models.Item{}

// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// Get single item by ID
func GetItemByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, item := range items {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    http.Error(w, "Item not found", http.StatusNotFound)
}

// Create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
    var newItem models.Item
    json.NewDecoder(r.Body).Decode(&newItem)
    newItem.ID = len(items) + 1
    items = append(items, newItem)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newItem)
}

// Delete an item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, item := range items {
        if item.ID == id {
            items = append(items[:index], items[index+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "Item not found", http.StatusNotFound)
}
