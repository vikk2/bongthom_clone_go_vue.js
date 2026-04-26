package data

import "backend/models"

type OrganizationTypeData interface {
	GetAll() []models.OrganizationType
}

type organizationTypeData struct {
	items []models.OrganizationType
}

func InitOrganizationTypeData() OrganizationTypeData {
	return &organizationTypeData{
		items: []models.OrganizationType{
			{ID: 1, Type: "Bar / Restaurant", TypeNum: 15},
			{ID: 2, Type: "Embassy", TypeNum: 1},
			{ID: 3, Type: "Government", TypeNum: 2},
			{ID: 4, Type: "Govt. NGO Project", TypeNum: 78},
			{ID: 5, Type: "Individual Person", TypeNum: 1},
			{ID: 6, Type: "International Org.", TypeNum: 78},
			{ID: 7, Type: "NGO", TypeNum: 76},
			{ID: 8, Type: "Other Small Company", TypeNum: 1},
			{ID: 9, Type: "Private Organization", TypeNum: 307},
			{ID: 10, Type: "Real Estate Agent", TypeNum: 2},
			{ID: 11, Type: "School / University", TypeNum: 15},
		},
	}
}

func (d *organizationTypeData) GetAll() []models.OrganizationType{
	return d.items
}
