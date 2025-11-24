package handlers

import (
"encoding/json"
"net/http"
"order-service/database"
"order-service/models"
"github.com/go-chi/chi/v5"
)

type CreateOrderRequest struct {
	UserID uint `json:"user_id"`
	Items  []struct {
		MenuItemID uint `json:"menu_item_id"`
		Quantity   int  `json:"quantity"`
	} `json:"items"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := models.Order{
		UserID: req.UserID,
	}

	var totalPrice float64
	for _, item := range req.Items {
		price := 2.50
		orderItem := models.OrderItem{
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
			Price:      price * float64(item.Quantity),
		}
		order.Items = append(order.Items, orderItem)
		totalPrice += orderItem.Price
	}

	order.TotalPrice = totalPrice

	if err := database.DB.Create(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	if err := database.DB.Preload("Items").Find(&orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
