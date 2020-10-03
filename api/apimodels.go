package api

type SignUp struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate int    `json:"age"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
