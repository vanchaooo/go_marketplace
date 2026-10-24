package repository

import (

)

type HistoryRepository struct {
	records map[int64][]string
}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{
		records: make(map[int64][]string),
	}
}

func (h HistoryRepository) Add(userID int64, record string) {
	h.records[userID] = append(h.records[userID], record)
}

func (h HistoryRepository) Get(userID int64) []string {
	history := h.records[userID]

	result := make([]string, len(history))
	copy(history, result)
	return result
}