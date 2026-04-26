package data

import "backend/models"

type OtherServiceData interface {
	GetAll() []models.OtherService
}

type otherServiceData struct{
	items []models.OtherService
}

func InitOtherServiceData() OtherServiceData{
	return &otherServiceData{
		items: []models.OtherService{
			{ID: 1, Service: "Khmer / English /French Phrasebook"},
		},
	}
}

func (d *otherServiceData) GetAll() []models.OtherService{
	return d.items
}