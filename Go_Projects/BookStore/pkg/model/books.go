package model

type Books struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"Author"`
	Price int `json:"Price"`
	CreatedAt string `json:"created_at"`
}