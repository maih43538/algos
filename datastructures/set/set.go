package set

import (
	"errors"
	"fmt"
)

var (
	ErrKeyAlreadyExists = errors.New("error:key_already_exists")
	ErrElementNotFound  = errors.New("error:element_not_found")
)

type Element struct {
	key   string
	value string
	next  *Element
}

type HashTable []*Element

func New(size int) HashTable {
	if size <= 0 {
		size = 100
	}
	return make(HashTable, size)
}

func (h HashTable) BucketCount() int {
	return len(h)
}

func (h HashTable) ContainsKey(key string) bool {
	// generate the index
	index, err := KeyToIndex(key, h.BucketCount())
	if err != nil {
		fmt.Println("is_key_duplicate:error:get_key:", err)
		return false
	}

	// traverse the linked list at h[index]
	var current *Element
	current = h[index]
	for {
		if current == nil {
			break
		}
		if current.key == key {
			// compare element's key with incoming key
			return true
		}
		current = current.next
	}

	return false
}

func (h HashTable) Insert(key, val string) error {
	// check if key is duplicate
	if h.ContainsKey(key) {
		return ErrKeyAlreadyExists
	}

	index, err := KeyToIndex(key, h.BucketCount())
	if err != nil {
		return err
	}

	n := h[index]
	h[index] = &Element{
		key:   key,
		value: val,
		next:  n,
	}

	// if table is reaching optimal capacity
	//if len(h) >= (.75 * 100) {
	//	err := h.Resize(size)
	//	return err
	//}
	return nil
}

func (h HashTable) FindByKey(key string) (string, error) {
	index, err := KeyToIndex(key, h.BucketCount())
	if err != nil {
		return "", err
	}
	item := h[index]
	if item != nil {
		if item.key == key {
			return item.value, nil
		}
	}
	return "", ErrElementNotFound
}

func (h HashTable) Delete(key string) error {
	// key to index
	index, err := KeyToIndex(key, h.BucketCount())
	if err != nil {
		return err
	}

	var current *Element
	current = h[index]

	// there are more elements
	if current.next != nil {
		// traverse the linked list at h[index]
		for {
			if current == nil {
				break
			}
			// compare element next's key values
			if current.next.key == key {
				// set element next equal to element next's next (un-linking the target element)
				// current.next = current.next.next
				h[index].next = current.next.next

				// if table is below size/4
				//if len(h) >= (.75 * 100) {
				//	err := h.Resize(size)
				//	return err
				//}

				return nil
			}
			current = current.next
		}
	}

	// there is only one element
	h[index] = nil // set it to nil

	// if table is below size/4
	//if len(h) >= (.75 * 100) {
	//	err := h.Resize(size)
	//	return err
	//}

	return nil
}

func (h HashTable) Resize(size int) error {
	// create a resized slice
	resized := New(size)

	// iterate through the original table and copy items over
	for _, bucket := range h {
		if bucket != nil {
			if bucket.next != nil {
				//var current *Element
				//current = bucket
				//for {
				//	// traverse linked list
				//	if current == nil {
				//		break
				//	}
				//	if h.ContainsKey(current.key) {
				//		// skip element since it already exists in new table
				//		continue
				//	}
				//	// insert element
				//	err := resized.Insert(current.key, current.value)
				//	if err != nil {
				//		return err
				//	}
				//	// reset current
				//	current = current.next
				//}

			}

			if resized.ContainsKey(bucket.key) {
				// skip element since it already exists in new table
				continue
			}

			// insert element
			err := resized.Insert(bucket.key, bucket.value)
			if err != nil {
				return err
			}
		}
	}

	h = resized

	fmt.Println("\n ", h.BucketCount())
	return nil
}
