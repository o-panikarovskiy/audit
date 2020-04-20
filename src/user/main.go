package user

// NewUserStore user service
func NewUserStore(rep IRepository) IService {
	return NewUserService(rep)
}

// ShutDown allows grasefull exit
func ShutDown(srv IService) {
	srv.ShutDown()
}
