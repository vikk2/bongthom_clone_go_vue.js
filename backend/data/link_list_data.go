package data

import "backend/models"

var LinkListItems = []models.Link{
	{ID: 1, Label: "Contact us", GreenIcon: false, New: false},
	{ID: 2, Label: "Post your CV", GreenIcon: false, New: false},
	{ID: 3, Label: "Add New Job", GreenIcon: false, New: false},
	{ID: 4, Label: "Add New Classified", GreenIcon: false, New: true},
	{ID: 5, Label: "Live Music & DJ Gig Guide", GreenIcon: true, New: false},
}
