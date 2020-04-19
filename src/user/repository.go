package user

// IReader interface
type IReader interface {
	Find(string) (*User, error)
	FindByEmail(string) (*User, error)
	FindAll() ([]*User, error)
}

// IWriter interface
type IWriter interface {
	Store(*User) (*User, error)
	Update(*User) (*User, error)
}

//IRepository repository interface
type IRepository interface {
	IReader
	IWriter
	ShutDown()
}
