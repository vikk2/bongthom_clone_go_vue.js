package data

import "backend/models"

type RssFileData interface {
	GetAll() []models.RssFile
}

type rssFileData struct {
	items []models.RssFile
}

func InitRssFileData() RssFileData {
	return &rssFileData{
		items: []models.RssFile{
			{ID: 1, Rss: "All current jobs"},
			{ID: 2, Rss: "All current classified"},
			{ID: 3, Rss: "BUSINESS / IT - For Sale"},
			{ID: 4, Rss: "BUSINESS / IT - For Sale"},
			{ID: 5, Rss: "COMPUTERS / IT - For Sale"},
			{ID: 6, Rss: "HOUSE/FLAT - For Rent"},
			{ID: 7, Rss: "LAND - For Sale"},
			{ID: 8, Rss: "SHOPS & SERVICES"},
			{ID: 9, Rss: "THINGS TO DO"},
			{ID: 10, Rss: "VEHICLES - For Sale"},
		},
	}
}

func (d *rssFileData) GetAll() []models.RssFile{
	return d.items
}
