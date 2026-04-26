package data

import "backend/models"

type SpecialScheduleData interface {
	GetAll() []models.SpecialSchedule
}

type specialScheduleData struct {
	items []models.SpecialSchedule
}

func InitSpecialScheduleData() SpecialScheduleData {
	return &specialScheduleData{
		items: []models.SpecialSchedule{
			{ID: 1, Sched_title: "Part-time", Sched_num: 14},
			{ID: 2, Sched_title: "Short-term", Sched_num: 4},
		},
	}
}

func (d *specialScheduleData) GetAll() []models.SpecialSchedule{
	return d.items
}
