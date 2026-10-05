package main
 
import (
	"sort"
	"strings"
)
 
func sortIDs(ids []int64) {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
}
 
func (m *Marketplace) GetUser(id int64) (User, bool) {
	user, exists := m.users[id]
	if !exists {
		return User{}, false
	}
 
	return *user, true
}
 
func (m *Marketplace) GetBalance(userID int64) (int64, bool) {
	wallet, exists := m.wallets[userID]
	if !exists {
		return 0, false
	}
 
	return wallet.Balance, true
}
 
func (m *Marketplace) GetProduct(id int64) (Product, bool) {
	product, exists := m.products[id]
	if !exists {
		return Product{}, false
	}
 
	return *product, true
}
 
func (m *Marketplace) GetCart(userID int64) (Cart, bool) {
	cart, ok := m.userCart(userID)
	if !ok {
		return Cart{}, false
	}
 
	itemsCopy := make(map[int64]int64, len(cart.Items))
	for productID, quantity := range cart.Items {
		itemsCopy[productID] = quantity
	}
 
	return Cart{UserID: cart.UserID, Items: itemsCopy}, true
}
 
func (m *Marketplace) GetOrder(userID, orderID int64) (Order, bool) {
	_, exists := m.users[userID]
	if !exists {
		return Order{}, false
	}
 
	order, exists := m.orders[orderID]
	if !exists || order.UserID != userID {
		return Order{}, false
	}
 
	orderCopy := *order
	orderCopy.Product = CopyOrderItems(order.Product)
	return orderCopy, true
}
 
func (m *Marketplace) GetOrderItems(userID, orderID int64) ([]OrderItem, bool) {
	order, ok := m.GetOrder(userID, orderID)
	if !ok {
		return nil, false
	}
 
	return order.Product, true
}

func (m *Marketplace) GetUserOrders(userID int64) []int64 {
	result := []int64{}
 
	_, exists := m.users[userID]
	if !exists {
		return result
	}
 
	for id, order := range m.orders {
		if order.UserID == userID {
			result = append(result, id)
		}
	}
 
	sortIDs(result)
	return result
}

func (m *Marketplace) GetUserHistory(userID int64) []string {
	history := m.operationHistory[userID]
 
	historyCopy := make([]string, len(history))
	copy(historyCopy, history)
	return historyCopy
}

func (m *Marketplace) FindUsersByName(query string) []int64 {
	query = strings.ToLower(query)
	result := []int64{}
 
	for id, user := range m.users {
		if strings.Contains(strings.ToLower(user.Name), query) {
			result = append(result, id)
		}
	}
 
	sortIDs(result)
	return result
}
 
func (m *Marketplace) SearchProducts(query string) []int64 {
	query = strings.ToLower(query)
	result := []int64{}
 
	for id, product := range m.products {
		if strings.Contains(strings.ToLower(product.Name), query) {
			result = append(result, id)
		}
	}
 
	sortIDs(result)
	return result
}