package sessions

// IStorage interface
type IStorage interface {
	Get(string) (string, error)
	Set(string, string, int) error
	Delete(string) error
}
