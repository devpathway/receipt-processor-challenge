package main

import (
	"sync"
)

var (
	store = make(map[string]*Receipt)
	lock  = sync.RWMutex{}
)

// SaveReceipt stores a receipt in memory
func SaveReceipt(id string, receipt *Receipt) {
	lock.Lock()
	defer lock.Unlock()
	store[id] = receipt
}

// GetReceipt retrieves a receipt by its ID
func GetReceipt(id string) (*Receipt, bool) {
	lock.RLock()
	defer lock.RUnlock()
	receipt, exists := store[id]
	return receipt, exists
}
