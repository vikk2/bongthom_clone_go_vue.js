package data

import "backend/models"


type AnnouncementData interface{
	GetAll() []models.Announcement
}

type announcementData struct{
	items []models.Announcement
}

func InitAnnouncementData() AnnouncementData{
	return &announcementData{
		items: []models.Announcement{
			{ID: 1, Type: "Call for Proposals", AnnouncementNum: 2 },
			{ID: 2, Type: "Expression of Interest", AnnouncementNum: 1},
			{ID: 3, Type: "Invitation for Bids", AnnouncementNum: 2},
			{ID: 4, Type: "Traning / Workshops", AnnouncementNum: 71},
		},
	}
}

func (d *announcementData) GetAll() []models.Announcement{
	return d.items
}
