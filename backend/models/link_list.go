package models

type Link struct {
	ID int `json:"id"`
	Label string `json:"link_item"`
	GreenIcon bool `json:"green_icon"`
	New bool `json:"new_label"`
}
