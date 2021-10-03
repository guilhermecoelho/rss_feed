package models

type Links struct {
	Links []Link `json:"links"`
}
type Link struct {
	Url string `json:"url"`
}
