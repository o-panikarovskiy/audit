package sessions

// IStorage interface
type IStorage interface {
	// Has returns false and empty error if value does not exist.
	Has(string) (bool, error)

	// Get returns empty string and empty error if value does not exist.
	Get(string) (string, error)

	//Set value by key and expiration in seconds
	Set(string, string, int) error

	// Delete returns false and empty error if valuse does not exist.
	Delete(string) (bool, error)

	// GetJSON returns error if value does not exist.
	GetJSON(string) (*map[string]interface{}, error)
}
