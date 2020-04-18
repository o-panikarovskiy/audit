package user

// IReader interface
type IReader interface {
	Find(id string) (*User, error)
	FindAll() ([]*User, error)
}

// IWriter interface
type IWriter interface {
	Save(user *User) (*User, error)
}

//IRepository repository interface
type IRepository interface {
	IReader
	IWriter
	ShutDown()
}
