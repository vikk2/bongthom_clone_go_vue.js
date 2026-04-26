package data

import "backend/models"

type LinkListData interface{
	GetAll() []models.Link
}

type linkListData struct{
	items []models.Link
}

func InitLinkListData() LinkListData{
	return &linkListData{
		items: []models.Link{
			{ID: 1, Label: "Contact us", GreenIcon: false, New: false},
			{ID: 2, Label: "Post your CV", GreenIcon: false, New: false},
			{ID: 3, Label: "Add New Job", GreenIcon: false, New: false},
			{ID: 4, Label: "Add New Classified", GreenIcon: false, New: true},
			{ID: 5, Label: "Live Music & DJ Gig Guide", GreenIcon: true, New: false},
		},
	}
}

func (d *linkListData) GetAll() []models.Link{
	return d.items
}
