package main

func (m *Marketplace) userCart(userID int64) (*Cart, bool) {
	_, exists := m.users[userID]
	if !exists {
		return nil, false
	}

	cart, ok := m.carts[userID]
	if !ok {
		return nil, false
	}

	return cart, true
}

func (m *Marketplace) AddToCart(userID, productID, quantity int64) bool {
	cart, ok := m.userCart(userID)
	if !ok {
		return false
	}

	_, exists := m.products[productID]
	if !exists {
		return false
	}

	return cart.Add(productID, quantity)
}

func (m *Marketplace) SetCartQuantity(userID, productID, quantity int64) bool {
	cart, ok := m.userCart(userID)
	if !ok {
		return false
	}
	_, exists := m.products[productID]
	if !exists {
		return false
	}
	_, inCart := cart.Items[productID]
	if !inCart {
		return false
	}

	return cart.SetQuantity(productID, quantity)
}

func (m *Marketplace) RemoveFromCart(userID, productID int64) bool {
	cart, ok := m.userCart(userID)
	if !ok {
		return false
	}

	return cart.Remove(productID)
}

func (m *Marketplace) ClearCart(userID int64) bool {
	cart, ok := m.userCart(userID)
	if !ok {
		return false
	}

	cart.Clear()
	return true
}

func (m *Marketplace) CalculateCartTotal(userID int64) (int64, bool) {
	cart, ok := m.userCart(userID)
	if !ok {
		return 0, false
	}

	items := make([]OrderItem, 0, len(cart.Items))
	for productID, quantity := range cart.Items {
		product, exists := m.products[productID]
		if !exists {
			return 0, false
		}

		item, ok := NewOrderItem(*product, quantity)
		if !ok {
			return 0, false
		}
		items = append(items, item)
	}

	return CalculateOrderTotal(items)
}