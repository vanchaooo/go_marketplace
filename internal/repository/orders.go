package repository

import (
	"math"
	"sort"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type OrderRepository struct {
	orders map[int64]*model.Order
	nextID int64
}

func NewOrderPerository() *OrderRepository{
	return &OrderRepository{
		orders: make(map[int64]*model.Order),
		nextID: 1,
	}
}

func (o OrderRepository) Add(order *model.Order) (int64, bool) {
	if o.nextID == 0 || o.nextID == math.MaxInt64 {
		return 0, false
	}

	order.OrderID = o.nextID
	o.orders[order.OrderID] = order
	o.nextID++
	return order.OrderID, true
}

func (o OrderRepository) Get(orderID int64) (*model.Order, bool) {
	order, ok := o.orders[orderID]
	return order, ok
}

func (o OrderRepository) GetAll(userID int64) []*model.Order {
	result := []*model.Order{}

	for _, order := range o.orders {
		if order.UserID == userID {
			result = append(result, order)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].OrderID < result[j].OrderID
	})
	return result
}