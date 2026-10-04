package main

type Marketplace struct {
	users            map[int64]*User
	wallets          map[int64]*Wallet
	products         map[int64]*Product
	carts            map[int64]*Cart
	orders           map[int64]*Order
	operationHistory map[int64][]string
	usedUserIDs      map[int64]bool
	nextOrderID      int64
}

func NewMarketplace() *Marketplace {
	return &Marketplace{
		users:            make(map[int64]*User),
		wallets:          make(map[int64]*Wallet),
		products:         make(map[int64]*Product),
		carts:            make(map[int64]*Cart),
		orders:           make(map[int64]*Order),
		operationHistory: make(map[int64][]string),
		usedUserIDs:      make(map[int64]bool),
		nextOrderID:      1,
	}
}