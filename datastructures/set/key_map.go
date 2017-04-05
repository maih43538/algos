package set

import "hash/fnv"

func KeyToIndex(key string, buckets int) (int, error) {
	hashedKey, err := GenerateHash(key)
	if err != nil {
		return 0, err
	}
	// modulo reduction to generate a number between size of slice
	index := hashedKey % buckets
	return index, nil
}

func GenerateHash(key string) (int, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(key))
	if err != nil {
		return 0, err
	}
	hashed := h.Sum32()
	return int(hashed), nil
}
