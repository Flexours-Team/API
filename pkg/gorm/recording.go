package gorm

import (
	"github.com/Flexours-Team/models"
	uuid "github.com/go-openapi/strfmt"
)

func (m *Manager) CreateStartScheduling(schedule *models.Schedule) (err error){
	err = m.DB.Create(&schedule).Error
	return
}

func (m *Manager) CreateEndScheduling( schedule *models.Schedule ) (err error) {
	err = m.DB.Model(&schedule).Updates(&schedule).Error
	if err != nil {
		return
	}
	var getSchedule *models.Schedule
	err = m.DB.First(&getSchedule, schedule.ID).Error
	if err != nil {
		return
	}

	//var newSchedule *models.Schedule
	//getSchedule.EndDate - getSchedule.StartDate)
	//fmt.Println(getSchedule.EndDate.Sub(getSchedule.StartDate).Seconds())
	schedule.HoursTotal = getSchedule.EndDate.Sub(getSchedule.StartDate).Seconds()
	err = m.DB.Model(&schedule).Updates(&schedule).Error

	//
	//var getSchedule2 *models.Schedule
	//err = m.DB.First(&getSchedule2, schedule.ID).Error
	//if err != nil {
	//	return
	//}
	//fmt.Println("HEREE: ",getSchedule2.HoursTotal.Seconds())


	return
}

func (m *Manager) GetSchedulings() (schedulings []models.Schedule, err error) {
	err = m.DB.Find(&schedulings).Order("startDate desc").Error
	//fmt.Println(schedulings)
	//fmt.Println()
	return
}

func (m *Manager) GetSchedulingByID(scheduleID uuid.UUID4) (schedule models.Schedule, err error) {
	err = m.DB.First(&schedule, scheduleID).Error
	return
}
func (m *Manager) GetDayHours (day1,day2 string) (hours float64, err error) {
	err = m.DB.Raw("SELECT sum(hours_total)  FROM \"schedules\" WHERE start_date between date('2021-04-"+day1+" 00:00:00') and date('2021-04-"+day2+" 00:00:00')").Scan(&hours).Error
	return
}