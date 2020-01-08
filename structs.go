package main

type LoginUser struct{
	Username  string `json:"username"`
	Password  string `json:"password"`
}
type Userdetails struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}
type ErrorResponse struct {
	Code    int
	Message string
}
type Login_token struct {
	Username  string `json:"username"`
	Token     string `json:"token"`
}