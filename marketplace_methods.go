package main

func (m *Marketplace) AddUser(userID int64, name string) bool {
	if m.usedUserIDs[userID] {
		return false
	}

	user, ok := NewUser(userID, name)
	if !ok {
		return false
	}
	wallet, ok := NewWallet(userID)
	if !ok {
		return false
	}
	cart, ok := NewCart(userID)
	if !ok {
		return false
	}
 
	m.users[userID] = &user
	m.wallets[userID] = &wallet
	m.carts[userID] = &cart
	m.usedUserIDs[userID] = true
	return true
}
 
func (m *Marketplace) RenameUser(userID int64, name string) bool {
	user, exists := m.users[userID]
	if !exists {
		return false
	}
 
	name, ok := normalizeName(name)
	if !ok {
		return false
	}
 
	user.Name = name
	return true
}
 
func (m *Marketplace) DeleteUser(userID int64) bool {
	_, exists := m.users[userID]
	if !exists {
		return false
	}
 
	wallet, exists := m.wallets[userID]
	if !exists || wallet.Balance != 0 {
		return false
	}
 
	for _, order := range m.orders {
		if order.UserID == userID && order.IsPaid() {
			return false
		}
	}
 
	delete(m.users, userID)
	delete(m.wallets, userID)
	delete(m.carts, userID)
	return true
}
 
func (m *Marketplace) TopUpBalance(userID, amount int64) bool {
	_, exists := m.users[userID]
	if !exists {
		return false
	}
 
	wallet, exists := m.wallets[userID]
	if !exists {
		return false
	}
 
	return wallet.TopUp(amount)
}
 
func (m *Marketplace) AddProduct(id int64, name string, price, stock int64) bool {
	_, exists := m.products[id]
	if exists {
		return false
	}
 
	product, ok := NewProduct(id, name, price, stock)
	if !ok {
		return false
	}
 
	m.products[id] = &product
	return true
}

func (m *Marketplace) UpdateProductStock(id, stock int64) bool {
	product, exists := m.products[id]
	if !exists {
		return false
	}
	if stock < 0 {
		return false
	}
 
	product.Stock = stock
	return true
}
 
func (m *Marketplace) RenameProduct(productID int64, name string) bool {
	product, exists := m.products[productID]
	if !exists {
		return false
	}
 
	name, ok := normalizeName(name)
	if !ok {
		return false
	}
 
	product.Name = name
	return true
}
 
func (m *Marketplace) UpdateProductPrice(productID, price int64) bool {
	product, exists := m.products[productID]
	if !exists {
		return false
	}
	if price <= 0 {
		return false
	}
 
	product.Price = price
	return true
}
 
func (m *Marketplace) AddProductStock(productID, quantity int64) bool {
	product, exists := m.products[productID]
	if !exists {
		return false
	}
 
	return product.AddStock(quantity)
}