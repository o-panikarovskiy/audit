package user

var userRepo IRepository
var userService IService

// Init db coonection and services
func Init() {
	userRepo = NewPostgreSQLRepository("some pool")
	userService = NewUserService(userRepo)
}

// GetService returns Service that implements business logic
func GetService() IService {
	return userService
}

// ShutDown allows grasefull exit
func ShutDown() {

}
