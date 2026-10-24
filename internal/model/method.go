package model



// User

func (u *User) Rename(name string) bool {
	name, ok := normalizeName(name)
	if !ok {
		return false
	}

	u.Name = name
	return true
}



// Wallet

func (w *Wallet) HasEnoughMoney(amount int64) bool {
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



// Product

func (p Product) IsAvailable(quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	return p.Stock >= quantity
}

func (p *Product) Rename(name string) bool {
	name, ok := normalizeName(name)
	if !ok {
		return false
	}

	p.Name = name
	return true
}

func (p *Product) SetPrice(price int64) bool {
	if price <= 0 {
		return false
	}

	p.Price = price
	return true
}

func (p *Product) SetStock(stock int64) bool {
	if stock < 0 {
		return false
	}

	p.Stock = stock
	return true
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



// Cart

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

func (c Cart) Copy() Cart {
	itemsCopy := make(map[int64]int64, len(c.Items))
	for productID, quantity := range c.Items {
		itemsCopy[productID] = quantity
	}

	return Cart{UserID: c.UserID, Items: itemsCopy}
}



// Order

func (o Order) IsPaid() bool {
	return o.Status == StatusPaid
}

func (o Order) IsCancelled() bool {
	return o.Status == StatusCancelled
}

func (o *Order) MarkCancelled() bool {
	if !o.IsPaid() {
		return false
	}

	o.Status = StatusCancelled
	return true
}

func (o Order) Copy() Order {
	o.Product = CopyOrderItems(o.Product)
	return o
}