package sessions

// IStorage interface
type IStorage interface {
	// Has returns false and empty error if value does not exist.
	Has(key string) (bool, error)

	// Get returns empty string and empty error if value does not exist.
	Get(key string) (string, error)

	// Set value by key and expiration in seconds
	Set(key string, value string, exp int) error

	// Delete returns false and empty error if valuse does not exist.
	Delete(key string) (bool, error)

	// GetJSON returns nil if value does not exist.
	GetJSON(key string) (*map[string]interface{}, error)

	// SetJSON set data by key and expiration in seconds
	SetJSON(key string, data interface{}, exp int) error
}
