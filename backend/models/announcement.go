package models

type Announcement struct {
	ID int `json:"id"`
	Type string `json:"type"`
	AnnouncementNum int `json:"announcement_num"`
}