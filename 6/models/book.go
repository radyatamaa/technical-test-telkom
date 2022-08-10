package models

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type BookResponse struct {
	UserId string `json:"user_id"`
	Data interface{} `json:"data"`
}

