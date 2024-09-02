package request

type UserRequest struct {
    Username     string   
    Password    string   
	Name    string  
}

type LoginRequest struct {
    Username string
    Password string
}