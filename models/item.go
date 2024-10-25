// models/item.go

// @Summary Create item
// @Description Create a new item
// @Tags items
// @Accept json
// @Produce json
// @Param item body Item true "Item to create"
// @Success 201 {object} Item
// @Failure 400 {object} ErrorResponse
// @Router /items [post]
package models

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}
