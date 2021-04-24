package gorm

import (
	"github.com/Flexours-Team/models"
	uuid "github.com/go-openapi/strfmt"
)

func (m *Manager) CreateAbsense(schedule *models.Absense) (err error){
	err = m.DB.Create(&schedule).Error
	return
}

func (m *Manager) UpdateAbsenseByID( absense *models.Absense ) (err error) {
	err = m.DB.Model(&absense).Updates(&absense).Error
	if err != nil {
		return
	}
	//var getSchedule *models.Absense
	//err = m.DB.First(&getSchedule, schedule.ID).Error
	//if err != nil {
	//	return
	//}

	//var newSchedule *models.Schedule
	//getSchedule.EndDate - getSchedule.StartDate)
	//fmt.Println(getSchedule.EndDate.Sub(getSchedule.StartDate).Seconds())
	//schedule.HoursTotal = getSchedule.EndDate.Sub(getSchedule.StartDate)
	//err = m.DB.Model(&schedule).Updates(&schedule).Error

	//
	//var getSchedule2 *models.Schedule
	//err = m.DB.First(&getSchedule2, schedule.ID).Error
	//if err != nil {
	//	return
	//}
	//fmt.Println("HEREE: ",getSchedule2.HoursTotal.Seconds())


	return
}

func (m *Manager) GetAbsenses() (absenses []models.Absense, err error) {
	err = m.DB.Find(&absenses).Error
	return
}

func (m *Manager) GetAbsenseByID(absenseID uuid.UUID4) (absense models.Absense, err error) {
	err = m.DB.First(&absense, absenseID).Error
	return
}

