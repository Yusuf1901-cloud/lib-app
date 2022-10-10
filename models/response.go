package models

type Response struct {
	Id int `json:"id"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type GetBooksResponse struct {
	Data []Book `json:"data"`
}

type BookResponse struct {
	Book Book `json:"book"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
