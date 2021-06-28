package main

type Entry struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Cache struct {
	Entries map[int][]Entry
}

// @TODO accept max parameter
func (c *Cache) Add(orderID int, entry Entry) {
	if c.Entries[orderID] == nil {
		c.Entries[orderID] = make([]Entry, 0)
	}
	c.Entries[orderID] = append(c.Entries[orderID], entry)
}

func (c *Cache) Delete(orderID int, entry Entry) {
	c.Entries[orderID] = make([]Entry, 0)
}

func (c Cache) Retrieve(orderID int, entry Entry) []Entry {
	return c.Entries[orderID]
}

func NewCache() *Cache {
	return &Cache{}
}
