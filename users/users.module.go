package users

func UserModule() *Module {
	// Define routes for UserController
	userController := NewController("/users", []Route{
		{Method: "GET", Path: "/", Handler: getUsers},
		{Method: "POST", Path: "/", Handler: createUser},
	})

	// Create a module with UserController
	return NewModule([]*Controller{userController})
}
