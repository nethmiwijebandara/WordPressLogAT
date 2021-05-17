package model

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Register struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AccessLog struct {
	IPAddress string `json:"message"`
	DateTime  string `json:"datetime"`
	Request   string `json:"request"`
}

type ErrorLog struct {
	DateTime    string `json:"datetime"`
	WarningType string `json:"warningtype"`
	PID         string `json:"pid"`
	Error       string `json:"error"`
}
