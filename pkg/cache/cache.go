package cache

import (
	"sync"
)

type Cache struct {
	// Using a mutex to protect concurrent access since more
	// nodes may access the same collection simultaneously.
	mu sync.RWMutex

	Collections map[string]*Btree

	Mode string // Temporal
}

var (
	SimpleCache = Cache{
		Collections: make(map[string]*Btree),
		Mode:        "temporal",
		mu:          sync.RWMutex{},
	}
)

func CreateCache() *Cache {
	return &SimpleCache
}

func (sc *Cache) Get(collection, key string) (any, bool) {
	sc.mu.RLock()
	defer sc.mu.RUnlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		return nil, false
	}

	// TODO: Implement getEverything (GET c.*; return all keys and values in the collection)

	return bt.Get(key)
}

func (sc *Cache) Set(collection, key string, value any) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		bt = &Btree{}
		sc.Collections[collection] = bt
	}

	bt.Set(key, value)
}

func (sc *Cache) Delete(collection, key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		return
	}

	bt.Delete(key)
}

func (sc *Cache) Drop(collection string) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		return false
	}

	bt.Clear()

	delete(sc.Collections, collection)

	return true
}

func (sc *Cache) Truncate(collection string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		return
	}

	for _, node := range bt.ListAllNodes() {
		bt.RemoveNode(node.Key)
	}
}

func (sc *Cache) Update(collection, key string, value any) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	bt, ok := sc.Collections[collection]
	if !ok {
		return
	}

	bt.Set(key, value)
}
