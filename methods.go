package main

func(u User) HasEnoughMoney(amount int) bool {
	return u.Balance >= amount
}

func (u *User) TopUp(amount int) bool {
	if amount <= 0 {
		return false
	}
 
	u.Balance += amount
	return true
}

func (u *User) Withdraw(amount int) bool {
	if amount <= 0 {
		return false
	}
	if !u.HasEnoughMoney(amount) {
		return false
	}
 
	u.Balance -= amount
	return true
}

func (p Product) IsAvailable(quantity int) bool {
	return p.Stock >= quantity
}

func (p *Product) AddStock(quantity int) bool {
	if quantity <= 0 {
		return false
	}
 
	p.Stock += quantity
	return true
}

func (p *Product) Reserve(quantity int) bool {
	if quantity <= 0 {
		return false
	}
	if !p.IsAvailable(quantity) {
		return false
	}
 
	p.Stock -= quantity
	return true
}

func (c Cart) IsEmpty() bool {
	return len(c.Items) == 0
}

func (c *Cart) Add(productID int, quantity int) bool {
	if productID <= 0 || quantity <= 0 {
		return false
	}

	if c.Items == nil {
		c.Items = make(map[int]int)
	}
 
	c.Items[productID] += quantity
	return true
}

func (c *Cart) SetQuantity(productID int, quantity int) bool {
	if productID <= 0 || quantity < 0 {
		return false
	}
 
	if quantity == 0 {
		delete(c.Items, productID)
		return true
	}
 
	if c.Items == nil {
		c.Items = make(map[int]int)
	}
 
	c.Items[productID] = quantity
	return true
}
 

func (c *Cart) Remove(productID int) bool {
	_, exists := c.Items[productID]
	if !exists {
		return false
	}
 
	delete(c.Items, productID)
	return true
}

func (c *Cart) Clear() {
	c.Items = make(map[int]int)
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