package main

import "errors"

// FIFO cache
// If handling larger data set, consider using another data structure

type UsersCache struct {
	Store []string
	Limit int
}

func (c *UsersCache) Add(s string) error {
	if c.Limit > 100 {
		return errors.New("Reached max capacity")
	}
	c.Store = append(c.Store, s)
	c.Limit += 1
	return nil
}

func (c *UsersCache) Find(s string) (string, error) {
	for _, val := range c.Store {
		if s == val {
			return val, nil
		}
	}
	return "", errors.New("nothing")
}
