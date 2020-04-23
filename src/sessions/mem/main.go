package mem

import (
	"audit/src/sessions"
	"encoding/json"
	"fmt"
	"sync"
)

type memStorage struct {
	data sync.Map
}

// NewStorage create redis storage
func NewStorage() sessions.IStorage {
	return &memStorage{}
}

// Has returns false and empty error if value does not exist.
func (r *memStorage) Has(key string) (bool, error) {
	_, ok := r.data.Load(key)
	return ok, nil
}

// Get returns empty string and empty error if value does not exist.
func (r *memStorage) Get(key string) (string, error) {
	val, ok := r.data.Load(key)

	if !ok {
		return "", nil
	}

	return fmt.Sprint(val), nil
}

// GetJSON returns nil if value does not exist.
func (r *memStorage) GetJSON(key string) (*map[string]interface{}, error) {
	str, ok := r.data.Load(key)
	if !ok {
		return nil, nil
	}

	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(fmt.Sprint(str)), &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

// SetJSON set value by key and expiration in seconds
func (r *memStorage) SetJSON(key string, data interface{}, expiration int) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Set(key, string(bytes), expiration)
}

//Set value by key and expiration in seconds
func (r *memStorage) Set(key string, value string, expiration int) error {
	r.data.Store(key, value)
	return nil
}

// Delete returns false and empty error if valuse does not exist.
func (r *memStorage) Delete(key string) (bool, error) {
	val, _ := r.Get(key)
	r.data.Delete(key)

	if val == "" {
		return false, nil
	}

	return true, nil
}
