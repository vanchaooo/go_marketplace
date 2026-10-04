package main

func (w Wallet) HasEnoughMoney(amount int64) bool {
	if amount <= 0 {
		return false
	}
	return w.Balance >= amount
}
 
func (w *Wallet) TopUp(amount int64) bool {
	if amount <= 0 {
		return false
	}
 
	w.Balance += amount
	return true
}
 
func (w *Wallet) Withdraw(amount int64) bool {
	if !w.HasEnoughMoney(amount) {
		return false
	}
 
	w.Balance -= amount
	return true
}

func (p Product) IsAvailable(quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	return p.Stock >= quantity
}
 
func (p *Product) AddStock(quantity int64) bool {
	if quantity <= 0 {
		return false
	}

	p.Stock += quantity
	return true
}

func (p *Product) Reserve(quantity int64) bool {
	if !p.IsAvailable(quantity) {
		return false
	}
 
	p.Stock -= quantity
	return true
}
 
func (c Cart) IsEmpty() bool {
	return len(c.Items) == 0
}
 
func (c *Cart) Add(productID int64, quantity int64) bool {
	if productID <= 0 || quantity <= 0 {
		return false
	}
	if c.Items == nil {
		c.Items = make(map[int64]int64)
	}
 
	c.Items[productID] += quantity
	return true
}
 
func (c *Cart) SetQuantity(productID int64, quantity int64) bool {
	if productID <= 0 || quantity < 0 {
		return false
	}
 
	if quantity == 0 {
		delete(c.Items, productID)
		return true
	}
 
	if c.Items == nil {
		c.Items = make(map[int64]int64)
	}
	c.Items[productID] = quantity
	return true
}
 
func (c *Cart) Remove(productID int64) bool {
	_, exists := c.Items[productID]
	if !exists {
		return false
	}
 
	delete(c.Items, productID)
	return true
}
 
func (c *Cart) Clear() {
	c.Items = make(map[int64]int64)
}

func (o Order) IsPaid() bool {
	return o.Status == "paid"
}
 
func (o Order) IsCancelled() bool {
	return o.Status == "cancelled"
}

func (o *Order) MarkCancelled() bool {
	if !o.IsPaid() {
		return false
	}
 
	o.Status = "cancelled"
	return true
}